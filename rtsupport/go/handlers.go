package main

import (
	"log"
	"time"

	"github.com/mitchellh/mapstructure"
	r "github.com/rethinkdb/rethinkdb-go"
)

const (
	ChannelStop = iota
	UserStop
	MessageStop
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	go func() {
		err = r.Table("channel").
			Insert(channel).
			Exec(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()
}

func changeFeedHelper(cursor *r.Cursor, tableName string, send chan Message, stop chan bool) {
	result := make(chan r.ChangeResponse)
	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			result <- change
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				cursor.Close()
				return
			case change := <-result:
				if change.NewValue != nil && change.OldValue == nil {
					send <- Message{tableName + " add", change.NewValue}
					log.Printf("sent channel add %s\n", tableName)
				}
				if change.NewValue != nil && change.OldValue != nil {
					send <- Message{tableName + " edit", change.NewValue}
					log.Printf("sent channel edit %s\n", tableName)
				}
				if change.NewValue == nil && change.OldValue != nil {
					send <- Message{tableName + " remove", change.OldValue}
					log.Printf("sent channel remove %s\n", tableName)
				}

			}
		}
	}()
}

func subscribeChannel(client *Client, data interface{}) {
	log.Println("ws in subscribeChannel")
	stop := client.NewStopChannel(ChannelStop)
	cursor, err := r.Table("channel").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(client.session)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	changeFeedHelper(cursor, "channel", client.send, stop)
}

func unsubscribeChannel(client *Client, data interface{}) {
	client.StopForKey(ChannelStop)
}

func subscribeUser(client *Client, data interface{}) {
	log.Println("ws in subscribeUser")
	stop := client.NewStopChannel(UserStop)
	cursor, err := r.Table("user").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(client.session)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}

	changeFeedHelper(cursor, "user", client.send, stop)
}

func unsubscribeUser(client *Client, data interface{}) {
	client.StopForKey(UserStop)
}

func editUser(client *Client, data interface{}) {
	var user User
	err := mapstructure.Decode(data, &user)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	log.Printf("edit user %+v\n", user)
	client.userName = user.Name
	go func() {
		err = r.Table("user").
			Get(client.id).
			Update(user).
			Exec(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()
}

func subscribeMessage(client *Client, data interface{}) {
	log.Println("ws in subscribeMessage")
	val, ok := (data.(map[string]interface{}))["channelId"]
	if !ok {
		return
	}
	channelId, ok := val.(string)
	if !ok {
		return
	}

	stop := client.NewStopChannel(MessageStop)
	cursor, err := r.Table("message").
		OrderBy(r.OrderByOpts{Index: r.Desc("createtAt")}).
		Filter(r.Row.Field("channelId").Eq(channelId)).
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(client.session)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}

	changeFeedHelper(cursor, "message", client.send, stop)
}

func unsubscribeMessage(client *Client, data interface{}) {
	client.StopForKey(MessageStop)
}

func addMessage(client *Client, data interface{}) {
	var message MessageChat
	err := mapstructure.Decode(data, &message)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	go func() {
		message.CreatedAt = time.Now()
		message.Author = client.userName
		err = r.Table("message").
			Insert(message).
			Exec(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()
}
