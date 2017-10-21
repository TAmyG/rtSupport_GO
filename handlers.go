package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	r "gopkg.in/gorethink/gorethink.v3"
)

func addChannel(client *Client, data interface{}) {
	fmt.Println("channel add")
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	//TODO insert in Rethink
	go func() {

		err = r.Table("channel").Insert(channel).Exec(client.session)

		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()

}

func subscribeChannel(client *Client, data interface{}) {
	fmt.Println("channel subscribe")
	go func() {
		cursor, err := r.Table("channel").Changes(r.ChangesOpts{IncludeInitial: true}).Run(client.session)

		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}
		var change r.ChangeResponse
		for cursor.Next(&change) {
			if change.NewValue != nil && change.OldValue == nil {
				client.send <- Message{"channel add", change.NewValue}
				fmt.Println("se agregó nuevo cannal")
			}
		}
	}()
}
