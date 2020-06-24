package controllers

import (
	"github.com/astaxie/beego"
	"liveassembly/models"
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
	pushStruct := CreateChannel()
	ret := models.Resp{
		Code: 200,
		Msg:  "Add Stream-Pushing Success",
		Data: pushStruct,
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}
