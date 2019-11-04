package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"github.com/gorilla/mux"

	"example.com/basics/hello"
)

func websocketHandler(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(1 * time.Second)
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case _ = <-ticker.C:
					// Write message back to browser
					provider := hello.NewFileProvider("./file")
					quotes, _ := provider.Provide()
					index := rand.Intn(len(quotes))
					// send message
					if err := websocket.JSON.Send(conn, quotes[index]); err != nil {
						log.Println(err)
					}
				}
			}
		}()

		time.Sleep(30 * time.Second)
		ticker.Stop()
		done <- true
	}
}

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.
		Methods([]string{"GET"}...).
		Path("/hello1").
		Handler(hello.NewHandler1())
	router.
		Methods([]string{"GET"}...).
		Path("/hello2").
		Handler(hello.NewHandler2())
	router.
		Methods([]string{"GET"}...).
		Path("/hello3").
		Handler(hello.NewHandler4())
	router.
		Methods([]string{"GET"}...).
		Path("/hello4").
		Handler(websocket.Handler(websocketHandler))

	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	fmt.Println("listening on 127.0.0.1:8080...")
	if err := http.ListenAndServe(net.JoinHostPort("127.0.0.1", "8080"), router); err != nil {
		log.Fatal(err)
	}
}
