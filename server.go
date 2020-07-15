package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 支持跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		_    int
		data []byte
	)
	// Upgrade: websocket
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	// websocket Conn
	for {
		// Text, Binary
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()

}

func main() {
	http.HandleFunc("/ws", wsHandler)
	_ = http.ListenAndServe(":7777", nil)
}
