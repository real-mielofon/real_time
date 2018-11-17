package main

import (
	"time"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}


func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// msgType, msg, err := socket.ReadMessage()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&Message); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)
		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage := Message{"error", err}
				if err := socket.WriteJSON(otMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			subscribeChannel(inMessage.Data)
		}
		// fmt.Printf("msgType: %d, msg: %v\n", msgType, string(msg))
		// if err = socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}

func addChannel(data interface{}) error {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	fmt.Println("added channel")
	return nil
}

func subscribeChannel()  {
	//TODO: rethinkDB query
	for{
		time.Sleep(time.Second*1)
		message := Message{"channel add",
			Channel{"1","Software Support"}
		}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	} 
}

func main() {
	router := NewRouter()
	router.Handel("channel add", addChannel)
	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
