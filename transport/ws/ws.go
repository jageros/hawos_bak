/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    ws
 * @Date:    2021/6/9 4:40 下午
 * @package: ws
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/protos/pbf"
	"github.com/jageros/hawos/recover"
	"github.com/jageros/hawos/transport"
	http2 "github.com/jageros/hawos/transport/http"
	"net"
	"net/http"
	"sync"
)

type OnReadHandle func(uid string, rData []byte, writer ISession)
type DisconnectHandle func(uid string)

type Target struct {
	GroupId    string
	Uids       []string
	UnlessUids []string
}

func (t *Target) CopyPbTarget(tg *pbf.Target) *Target {
	t.GroupId = tg.GroupId
	t.Uids = tg.Uids
	t.UnlessUids = tg.UnlessUids
	return t
}

type Writer interface {
	Write(data []byte, target *Target) error
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func upgradeWsConn(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return upGrader.Upgrade(w, r, nil)
}

type Server struct {
	*transport.BaseServer
	svr          *http.Server
	sessions     map[string]*session
	readFn       []OnReadHandle
	disconnectFn DisconnectHandle

	rw *sync.RWMutex
}

func (s *Server) ConnCnt() int {
	return len(s.sessions)
}

func New(ctx context.Context, opfs ...transport.SvrOpFn) *Server {
	ss := &Server{
		BaseServer: transport.NewBaseServer(ctx, opfs...),
		sessions: map[string]*session{},
		rw:       &sync.RWMutex{},
	}
	ss.Options.Protocol = transport.WS

	addr := fmt.Sprintf("%s:%d", ss.Options.Ip, ss.Options.Port)
	ss.svr = &http.Server{
		Addr:    addr,
		Handler: ss,
		BaseContext: func(listener net.Listener) context.Context {
			return ss.Ctx
		},
	}

	return ss
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uid := r.Header.Get(http2.HTTP_HD_APP_UID)
	if uid == "" || uid == "undefined" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	conn, err := upgradeWsConn(w, r)
	if err != nil {
		log.Errorf("Websocket connect err=%v", err)
		return
	}
	sess := newSession(s.Ctx, uid, conn, s.Options.ReadTimeout, s.Options.WriteTimeout)
	s.rw.Lock()
	s.sessions[sess.uid] = sess
	s.rw.Unlock()
	g := &sync.WaitGroup{}
	g.Add(3)
	go sess.startHandleMsg(g, s.readFn...)
	go sess.startRecvMsg(g)
	go sess.startWriteMsg(g)
	g.Wait()
	sess.stop()
	s.rw.Lock()
	delete(s.sessions, sess.uid)
	s.rw.Unlock()
	if s.disconnectFn != nil {
		s.disconnectFn(sess.uid)
	}
}

func (s *Server) sendToGroup(data []byte, groupId string, unlessUids ...string) error {
	unless := map[string]bool{}
	for _, ulid := range unlessUids {
		unless[ulid] = true
	}

	return recover.CatchPanic(func() error {
		s.rw.RLock()
		gw := &sync.WaitGroup{}
		for _, sess := range s.sessions {
			if sess.groupId == groupId && !unless[sess.uid] {
				gw.Add(1)
				go func() {
					defer gw.Done()
					select {
					case <-sess.Ctx.Done():
						return
					case sess.writeMsgChan <- &writeMsg{
						messageType: websocket.BinaryMessage,
						data:        data,
					}:
					}
				}()
			}
		}
		s.rw.RUnlock()
		gw.Wait()
		return nil
	})
}

func (s *Server) sendUnless(data []byte, unlessUids ...string) error {
	unless := map[string]bool{}
	for _, ulid := range unlessUids {
		unless[ulid] = true
	}

	return recover.CatchPanic(func() error {
		s.rw.RLock()
		gw := &sync.WaitGroup{}
		for _, sess := range s.sessions {
			if !unless[sess.uid] {
				gw.Add(1)
				go func() {
					defer gw.Done()
					select {
					case <-sess.Ctx.Done():
						return
					case sess.writeMsgChan <- &writeMsg{
						messageType: websocket.BinaryMessage,
						data:        data,
					}:
					}
				}()
			}
		}
		s.rw.RUnlock()
		gw.Wait()
		return nil
	})
}

func (s *Server) sendToUsers(data []byte, uids ...string) error {
	return recover.CatchPanic(func() error {
		s.rw.RLock()
		gw := &sync.WaitGroup{}
		for _, sid := range uids {
			sess, ok := s.sessions[sid]
			if ok {
				gw.Add(1)
				go func() {
					defer gw.Done()
					select {
					case <-sess.Ctx.Done():
						return
					case sess.writeMsgChan <- &writeMsg{
						messageType: websocket.BinaryMessage,
						data:        data,
					}:
					}
				}()
			}
		}
		s.rw.RUnlock()
		gw.Wait()
		return nil
	})
}

func (s *Server) Write(data []byte, target *Target) error {
	if target.GroupId != "" {
		return s.sendToGroup(data, target.GroupId, target.UnlessUids...)
	} else if len(target.Uids) > 0 {
		return s.sendToUsers(data, target.Uids...)
	} else if len(target.UnlessUids) > 0 {
		return s.sendUnless(data, target.UnlessUids...)
	} else {
		return s.Broadcast(data)
	}
}

func (s *Server) Broadcast(data []byte) error {
	return recover.CatchPanic(func() error {
		s.rw.RLock()
		gw := &sync.WaitGroup{}
		for _, sess := range s.sessions {
			gw.Add(1)
			go func(sess *session) {
				defer gw.Done()
				select {
				case <-sess.Ctx.Done():
				case sess.writeMsgChan <- &writeMsg{
					messageType: websocket.BinaryMessage,
					data:        data,
				}:
				}
			}(sess)
		}
		s.rw.RUnlock()
		gw.Wait()
		return nil
	})
}

func (s *Server) RegistryReadFunc(onRead ...OnReadHandle) {
	s.readFn = onRead
}

func (s *Server) RegisterDisconnectFunc(disfn DisconnectHandle) {
	s.disconnectFn = disfn
}

func (s *Server) Serve() error {
	s.PrintInfo()
	return s.svr.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(s.Ctx, s.Options.CloseTimeout)
	defer cancel()
	s.rw.Lock()
	defer s.rw.Unlock()
	for _, ses := range s.sessions {
		ses.stop()
	}
	err := s.svr.Shutdown(ctx)
	if err != nil {
		log.Errorf("WebSocket Listen Shutdown err=%v", err)
	}
}
