package controllers

import (
	"github.com/agnaite/chatboat/app/consumer"
	"github.com/agnaite/chatboat/app/producer"
	"github.com/revel/revel"
)

func init() {
	go consumer.Consumer()
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Post(myMsg string) revel.Result {
	producer.Producer(myMsg)
	return c.Render(myMsg)
}
