package tlsSpoofer_go

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WSServer struct{}

func (ws *WSServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (ws *WSServer) Serve() error {
	config := NewConfig()
	s := &http.Server{
		Handler:     ws,
		ReadTimeout: time.Duration(config.ReadTimeout),
	}
	uprgader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return s.ListenAndServe()
}
