package controllers

import (
	"log"

	"github.com/revel/revel"
	"golang.org/x/net/websocket"
)

func init() {
}

type App struct {
	*revel.Controller
}

// var feed []string

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) EnterChat(user string) revel.Result {
	return c.Redirect("/websocket/room?user=%s", user)
}

func Publish(myMsg string, user string) {
	origin := "http://localhost/"
	url := "ws://localhost:6969/websocket/room/send?user=" + user + "&msg=" + myMsg
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := ws.Write([]byte(myMsg)); err != nil {
		log.Fatal(err)
	}
}
