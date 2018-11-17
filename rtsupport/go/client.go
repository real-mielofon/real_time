package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Client struct {
	send chan Message
}

func (client *Client) write() {
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
	}
}

func (client *Client) subscribeChannels() {
	for {
		time.Sleep(r())
		client.send <- Message{"channel add", ""}
	}
}

func (client *Client) subscribeMessages() {
	for {
		time.Sleep(r())
		client.send <- Message{"message add", ""}
	}
}

func r() time.Duration {
	return time.Microsecond * time.Duration(rand.Intn(1000))
}

func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}
