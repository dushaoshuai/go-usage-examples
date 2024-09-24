package quickstart

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
		if err := conn.WriteJSON(map[string]any{"type": messageType, "data": string(p)}); err != nil {
			log.Println(err)
			return
		}
	}
}

func run() {
	http.HandleFunc("/ws", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
