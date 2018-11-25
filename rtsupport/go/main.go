package main

import (
	"log"
	"net/http"
	"time"

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

type MessageChat struct {
	Id        string    `gorethink:"id,omitempty"`
	ChannelId string    `json:"channelId" gorethink:"channelId"`
	Author    string    `json:"author" gorethink:"author"`
	Body      string    `json:"body" gorethink:"body"`
	CreatedAt time.Time `json:"createdAt" gorethink:"createdAt"`
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

	err = r.Table("user").
		Filter(true).
		Delete().
		Exec(session)
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

	router.Handle("message add", addMessage)
	router.Handle("message subscribe", subscribeMessage)
	router.Handle("message unsubscribe", unsubscribeMessage)

	http.Handle("/", router)
	log.Println("http://localhost:4000 started!")
	http.ListenAndServe(":4000", nil)
}
