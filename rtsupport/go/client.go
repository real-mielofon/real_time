package main

import (
	"log"

	"github.com/gorilla/websocket"
	r "github.com/rethinkdb/rethinkdb-go"
)

type FindHandler func(string) (Handler, bool)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send         chan Message
	socket       *websocket.Conn
	findHandler  FindHandler
	session      *r.Session
	stopChannels map[int]chan bool
	id           string
	userName     string
}

func (c *Client) NewStopChannel(key int) chan bool {
	c.StopForKey(key)
	stop := make(chan bool)
	c.stopChannels[key] = stop
	return stop
}

func (c *Client) StopForKey(key int) {
	if ch, found := c.stopChannels[key]; found {
		ch <- true
		delete(c.stopChannels, key)
	}
}

func (c *Client) Read() {
	var message Message
	for {
		if err := c.socket.ReadJSON(&message); err != nil {
			break
		}
		log.Printf("ws: in %+v", message)
		if handler, found := c.findHandler(message.Name); found {
			log.Printf("ws: find handler")
			handler(c, message.Data)
		} else {
			log.Printf("ws: not find handler")
		}
	}
	c.socket.Close()
}

func (c *Client) Write() {
	for msg := range c.send {
		log.Printf("ws: out %+v", msg)
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

func (c *Client) Close() {
	for _, ch := range c.stopChannels {
		ch <- true
	}
	close(c.send)
}
func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	var user User
	user.Name = "anonymouse"
	res, err := r.Table("user").Insert(user).RunWrite(session)
	if err != nil {
		log.Println(err)
	}
	var id string
	if len(res.GeneratedKeys) > 0 {
		id = res.GeneratedKeys[0]
	}
	return &Client{
		send:         make(chan Message),
		socket:       socket,
		findHandler:  findHandler,
		session:      session,
		stopChannels: make(map[int]chan bool),
		id:           id,
		userName:     user.Name,
	}
}
