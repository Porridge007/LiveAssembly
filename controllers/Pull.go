package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"liveassembly/models"
)

type PullController struct {
	beego.Controller
}

func (c *PullController) Post() {
	pullAddr := c.GetString("pullAddr")
	fmt.Println(pullAddr)
	pushStruct := CreateChannel()

	go PullStream(pullAddr, pushStruct.PushAddr)
	ret := models.Resp{
		Code: 200,
		Msg:  "Pull Stream Success",
		Data: pushStruct,
	}

	c.Data["json"] = &ret
	c.ServeJSON()
}
