package controllers

import (
	"github.com/agnaite/chatboat/app/producer"
	"github.com/revel/revel"
	"net/http"
)

func init() {
}

type App struct {
	*revel.Controller
}

var feed []string

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
	feed = append(feed, myMsg)
}
