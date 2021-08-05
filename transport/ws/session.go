/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    session
 * @Date:    2021/6/9 6:06 下午
 * @package: ws
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/recover"
	"sync"
	"time"
)

const (
	readMsgChanLen  = 128
	writeMsgChanLen = 128
)

type ISession interface {
	EnterGroup(groupId string)
	Write(p []byte) (n int, err error)
}

type readMsg struct {
	messageType int
	data        []byte
}

type writeMsg struct {
	messageType int
	data        []byte
}

type session struct {
	uid          string
	groupId      string
	conn         *websocket.Conn
	readMsgChan  chan *readMsg
	writeMsgChan chan *writeMsg

	readTimeout  time.Duration
	writeTimeout time.Duration

	Ctx    context.Context
	cancel context.CancelFunc
}

func (s *session) Write(p []byte) (n int, err error) {
	err = recover.CatchPanic(func() error {
		select {
		case <-s.Ctx.Done():
			return s.Ctx.Err()
		case s.writeMsgChan <- &writeMsg{
			messageType: websocket.BinaryMessage,
			data:        p,
		}:
		}
		return nil
	})
	return len(p), err
}

func (s *session) EnterGroup(groupId string) {
	s.groupId = groupId
}

func (s *session) stop() {
	s.cancel()
	s.conn.Close()
}

func newSession(ctx_ context.Context, uid string, conn *websocket.Conn, readTimeout, writeTimeout time.Duration) *session {
	ctx, cancel := context.WithCancel(ctx_)
	ss := &session{
		uid:          uid,
		conn:         conn,
		readMsgChan:  make(chan *readMsg, readMsgChanLen),
		writeMsgChan: make(chan *writeMsg, writeMsgChanLen),
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
		Ctx:          ctx,
		cancel:       cancel,
	}
	return ss
}

func (s *session) startHandleMsg(g *sync.WaitGroup, handles ...OnReadHandle) {
	defer g.Done()
	defer s.stop()
	for {
		select {
		case msg := <-s.readMsgChan:
			g := &sync.WaitGroup{}
			for _, readHandle := range handles {
				g.Add(1)
				go func(sid string, handle OnReadHandle, data []byte) {
					defer g.Done()
					err := recover.CatchPanic(func() error {
						handle(sid, data, s)
						return nil
					})
					if err != nil {
						log.Errorf("Websocket msg handle err=%v", err)
					}
				}(s.uid, readHandle, msg.data)
			}
			g.Wait()

		case <-s.Ctx.Done():
			if len(s.readMsgChan) > 0 {
				continue
			}
			return
		}
	}
}

func (s *session) read() <-chan *readMsg {
	ch := make(chan *readMsg)
	go func() {
		var msg *readMsg // 出错的情况下 msg = nil
		err := s.conn.SetReadDeadline(time.Now().Add(s.readTimeout))
		if err == nil {
			mty, data, err := s.conn.ReadMessage()
			if err == nil {
				msg = &readMsg{
					messageType: mty,
					data:        data,
				}
			}
		}

		select {
		case <-s.Ctx.Done():
			return
		case ch <- msg:
		}
	}()
	return ch
}

func (s *session) startRecvMsg(g *sync.WaitGroup) {
	defer g.Done()
	for {
		select {
		case <-s.Ctx.Done():
			return
		case msg := <-s.read():
			if msg == nil {
				s.cancel()
				return
			}
			select {
			case <-s.Ctx.Done():
				return
			case s.readMsgChan <- msg:
			}
		}
	}
}

func (s *session) startWriteMsg(g *sync.WaitGroup) {
	defer g.Done()
	for {
		select {
		case <-s.Ctx.Done():
			return
		case msg := <-s.writeMsgChan:
			start := time.Now()
			err := s.conn.SetWriteDeadline(time.Now().Add(s.writeTimeout))
			if err != nil {
				log.Errorf("Websocket SetWriteDeadline err=%v", err)
				s.cancel()
				return
			}

			err = s.conn.WriteMessage(msg.messageType, msg.data)
			if err != nil {
				log.Errorf("Websocket WriteMessage err=%v", err)
				s.cancel()
				return
			}
			end := time.Now()
			take := end.Sub(start)
			if take > time.Second {
				log.Infof("Write Msg to Client uid=%s take %s", s.uid, take.String())
			}
		}
	}
}
