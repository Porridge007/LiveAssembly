package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"liveassembly/models"
	"os/exec"
	"strconv"
)

var RoomID int
var PullRoom = make(models.PullRoomMap)

type Channel struct {
	Status  int    `json:"status"`
	Channel string `json:"data"`
}

type PullController struct {
	beego.Controller
}

type PushController struct {
	beego.Controller
}

type StarterController struct {
	beego.Controller
}

type KillerController struct {
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

func (c *PullController) Post() {
	pullAddr := c.GetString("pullAddr")
	pushStruct := CreateChannel()

	PullRoom[GetChannel(pushStruct.WatchAddr.Rtmp)] = models.Addrs{
		PullAddr: pullAddr,
		PushAddr: pushStruct.PushAddr,
	}

	//go PullStream(pullAddr, pushStruct.PushAddr, GetChannel(pushStruct.WatchAddr.Rtmp))
	ret := models.Resp{
		Code: 200,
		Msg:  "Pull Stream Success",
		Data: pushStruct,
	}

	c.Data["json"] = &ret
	c.ServeJSON()
}

func (c *StarterController) Post() {
	RoomID := c.GetString("room")
	addrs := PullRoom[RoomID]
	fmt.Println(addrs)
	go PullStream(addrs.PullAddr, addrs.PushAddr, RoomID)
	ret := models.Resp{
		Code: 0,
		Msg:  "Start Stream-pulling success",
		Data: nil,
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

func (c *KillerController) Post() {
	RoomID := c.GetString("room")
	cmd := exec.Command("taskkill", "/pid", strconv.Itoa(PidMap[RoomID]), "/f")
	fmt.Println(cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	go cmd.Run()
	ret := models.Resp{
		Code: 0,
		Msg:  "Kill Stream-pulling success",
		Data: nil,
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}
