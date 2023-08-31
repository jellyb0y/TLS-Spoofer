package wsServer

import (
	"fmt"
	"github.com/gorilla/websocket"
	. "github.com/jellyb0y/TLS-Spoofer/tlsSpoofer-go/config"
	"log"
	"net/http"
	"time"
)

type WSServer struct {
	config *Config
}

func NewWSServer(options Options) *WSServer {
	return &WSServer{config: NewConfig(options)}
}

func (ws *WSServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  ws.config.ReadBufferSize,
		WriteBufferSize: ws.config.WriteBufferSize,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (ws *WSServer) Serve() error {
	s := &http.Server{
		Handler:     ws,
		ReadTimeout: time.Duration(ws.config.ReadTimeout),
		Addr:        fmt.Sprintf(":%d", ws.config.ListenPort),
	}
	http.Handle("/", ws)

	return s.ListenAndServe()
}
