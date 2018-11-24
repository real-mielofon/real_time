package main

import (
	"fmt"
	"net/http"

	r "github.com/rethinkdb/rethinkdb-go"
)

type Handler func(client *Client, data interface{})

type Router struct {
	rules   map[string]Handler
	session *r.Session
}

func (r *Router) FindHandler(msgName string) (handler Handler, found bool) {
	handler, found = r.rules[msgName]
	return handler, found
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(socket, r.FindHandler, r.session)
	defer client.Close()

	go client.Write()
	client.Read()
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

func NewRouter(session *r.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}
