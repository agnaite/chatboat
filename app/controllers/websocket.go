package controllers

import (
	"net/http"
	"os"

	"github.com/agnaite/chatboat/app/producer"
	"github.com/revel/examples/chat/app/chatroom"
	"github.com/revel/revel"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) Room(user string) revel.Result {
	api := os.Getenv("GIPHY_API")
	return c.Render(user, api)
}

func (c WebSocket) Post(w http.ResponseWriter, r *http.Request) revel.Result {
	myMsg := c.Params.Form.Get("msg")
	user := c.Params.Form.Get("user")

	producer.Producer(myMsg, user)
	return nil
}

func (c WebSocket) SendMsg(user string, msg string, ws revel.ServerWebSocket) revel.Result {
	chatroom.Say(user, msg)
	return nil
}

func (c WebSocket) RoomSocket(user string, ws revel.ServerWebSocket) revel.Result {
	// Make sure the websocket is valid.
	if ws == nil {
		return nil
	}

	// Join the room.
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user)
	defer chatroom.Leave(user)

	// Send down the archive.
	for _, event := range subscription.Archive {
		if ws.MessageSendJSON(&event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := ws.MessageReceiveJSON(&msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroom.
	for {
		select {
		case event := <-subscription.New:
			if ws.MessageSendJSON(&event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			chatroom.Say(user, msg)
		}
	}

	return nil
}
