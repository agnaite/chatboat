package controllers

import (
	"fmt"
	"github.com/agnaite/chatboat/app/producer"
	"github.com/revel/revel"
	//"time"
)

func init() {
}

type App struct {
	*revel.Controller
}

var feed []string

func (c App) Index() revel.Result {
	return c.Render(feed)
}

func (c App) Post(myMsg string) revel.Result {
	producer.Producer(myMsg)
	return c.Redirect(App.Index)
}

func Publish(myMsg string) {
	feed = append(feed, myMsg)
	fmt.Println("in PUBLISH ", myMsg)
}
