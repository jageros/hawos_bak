/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    Session
 * @Date:    2021/6/9 6:06 下午
 * @package: ws
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package wsc

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/jageros/hawos/internal/pkg/log"
	recover3 "github.com/jageros/hawos/recover"
	"net/http"
	"sync"
	"time"
)

const (
	readTimeout = time.Second * 30
	writeTime   = time.Second * 30
)

type readMsg struct {
	messageType int
	data        []byte
}

type writeMsg struct {
	messageType int
	data        []byte
}

type Session struct {
	Id           string
	conn         *websocket.Conn
	readMsgChan  chan *readMsg
	writeMsgChan chan *writeMsg

	readTimeout  time.Duration
	writeTimeout time.Duration

	Ctx    context.Context
	cancel context.CancelFunc

	Handle func(uid string, data []byte)
}

func (s *Session) RegistryHandle(f func(uid string, data []byte)) {
	s.Handle = f
}

func (s *Session) Write(p []byte) (n int, err error) {
	err = recover3.CatchPanic(func() error {
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

func (s *Session) Stop() {
	s.cancel()
	s.conn.Close()
}

func newSession(ctx context.Context, uid string, conn *websocket.Conn, readTimeout, writeTimeout time.Duration) *Session {
	ctx_, cancel := context.WithCancel(ctx)
	ss := &Session{
		Id:           uid,
		conn:         conn,
		readMsgChan:  make(chan *readMsg, 16),
		writeMsgChan: make(chan *writeMsg, 16),
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
		Ctx:          ctx_,
		cancel:       cancel,
		Handle: func(uid string, data []byte) {
			log.Infof("Recv msg uid=%s", uid)
		},
	}
	return ss
}

func (s *Session) startHandleMsg(g *sync.WaitGroup) {
	defer g.Done()
	for {
		select {
		case <-s.Ctx.Done():
			return
		case msg := <-s.readMsgChan:
			if s.Handle != nil {
				err := recover3.CatchPanic(func() error {
					s.Handle(s.Id, msg.data)
					return nil
				})
				if err != nil {
					log.Infof("Uid=%s websocket msg handle err=%v", s.Id, err)
				}
			}
		}
	}
}

func (s *Session) read() <-chan *readMsg {
	ch := make(chan *readMsg)
	go func() {
		recover3.CatchPanic(func() error {
			var msg *readMsg
			err := s.conn.SetReadDeadline(time.Now().Add(s.readTimeout))
			if err == nil {
				ty, p, err := s.conn.ReadMessage()
				if err == nil {
					msg = &readMsg{
						messageType: ty,
						data:        p,
					}
				}
			}

			select {
			case <-s.Ctx.Done():
				return s.Ctx.Err()
			case ch <- msg:
			}
			return nil
		})
	}()
	return ch
}

func (s *Session) startRecvMsg(g *sync.WaitGroup) {
	defer g.Done()
	for {
		select {
		case <-s.Ctx.Done():
			return
		case msg := <-s.read():
			if msg == nil {
				continue
			}
			select {
			case <-s.Ctx.Done():
				return
			case s.readMsgChan <- msg:
			}
		}
	}
}

func (s *Session) startWriteMsg(g *sync.WaitGroup) {
	defer g.Done()
	for {
		select {
		case <-s.Ctx.Done():
			return
		case msg := <-s.writeMsgChan:
			err := s.conn.SetWriteDeadline(time.Now().Add(s.writeTimeout))
			if err != nil {
				log.Infof("websocket SetWriteDeadline err=%v", err)
				s.cancel()
				return
			}

			err = s.conn.WriteMessage(msg.messageType, msg.data)
			if err != nil {
				log.Infof("websocket WriteMessage err=%v", err)
				s.cancel()
				return
			}
		}
	}
}

func (s *Session) Run() {
	g := &sync.WaitGroup{}
	g.Add(3)
	go s.startHandleMsg(g)
	go s.startRecvMsg(g)
	go s.startWriteMsg(g)
	g.Wait()
}

func Connect(ctx context.Context, uid string, url string, header http.Header) (*Session, error) {
	if header == nil {
		header = http.Header{}
	}
	header.Set("Origin", "http://localhost/")
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, url, header)
	if err != nil {
		return nil, err
	}
	sess := newSession(ctx, uid, conn, readTimeout, writeTime)
	return sess, nil
}
