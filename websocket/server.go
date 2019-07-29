package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"flag"
	"log"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandle(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("welcome"))
	conn, err := upgrader.Upgrade(w, r, nil)
	if  err != nil {
		log.Print("upgrader", err)
	}

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		pp := []byte("43")
		conn.WriteMessage(messageType, pp)
	}
}

func main() {
	flag.Parse()
    log.SetFlags(0)
	http.HandleFunc("/ws", wsHandle)

	http.ListenAndServe("0.0.0.0:8888", nil)
}