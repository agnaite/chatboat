package controllers

import (
	"log"
	"net/http"

	"github.com/agnaite/chatboat/app/producer"
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

func (c App) Post(w http.ResponseWriter, r *http.Request) revel.Result {
	params := r.URL.Query()
	myMsg := params.Get("msg")
	producer.Producer(myMsg)
	return c.Redirect("/")
}

func Publish(myMsg string) {
	origin := "http://localhost/"
	url := "ws://localhost:6969/websocket/room/send?user=uhh&msg=" + myMsg
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte(myMsg)); err != nil {
		log.Fatal(err)
	}
}
