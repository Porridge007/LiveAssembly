package controllers

import (
	"github.com/astaxie/beego"
)

var RoomID int

type Channel struct {
	Status  int    `json:"status"`
	Channel string `json:"data"`
}

type PushController struct {
	beego.Controller
}

func (c *PushController) Post() {
	ret := CreateChannel()
	c.Data["json"] = &ret
	c.ServeJSON()

	RoomID += 1
}
