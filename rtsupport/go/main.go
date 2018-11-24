package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	r "github.com/rethinkdb/rethinkdb-go"
)

type Channel struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		log.Panic(err.Error())
		return
	}

	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)

	router.Handle("user edit", editUser)
	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)
	/*
		router.Handle("message add", addMessage)
		router.Handle("message subscribe", subscribeMessage)
		router.Handle("message unsubscribe", unsubscribeMessage)
	*/

	http.Handle("/", router)
	log.Println("http://localhost:4000 started!")
	http.ListenAndServe(":4000", nil)
}
