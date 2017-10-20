package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Router struct {
	rules map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

func (e *Router) ServerHTTP(w http.ResponseWriter, r *http.ReadRequest) {
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	client := NewClient(socket)
	go client.Write()
	client.Read()
}

func addChannel(data interface{}) error {
	var channel Channel
	//channelmap := data.(map[string]interface{})
	//channel.Name = channelmap["name"].(string)
	err := mapstructure.Decode(data, &channel)

	if err != nil {
		fmt.Println(err)
		return err
	}
	channel.Id = "1"

	fmt.Println("added channel")
	return nil
}
func subscribeChannel(socket *websocket.Conn) {
	//TODO: rethinkDB Query
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add",
			Channel{"1", "Software support"}}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}
