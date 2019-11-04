package hello

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler3 struct{}

func NewHandler3() Handler3 {
	return Handler3{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h Handler3) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	conn, _ := upgrader.Upgrade(res, req, nil) // error ignored for sake of simplicity
	for {
		// Read message from browser
		msgType, _, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		if err = conn.WriteMessage(msgType, []byte("test")); err != nil {
			log.Fatal(err)
		}

	}

}
