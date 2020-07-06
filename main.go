package main

import (
	"github.com/astaxie/beego"
	_ "liveassembly/routers"
	"os/exec"
)

func init(){
	go execLiveGo()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func execLiveGo(){
	cmd:= exec.Command("taskkill /f /t /im livego.exe")
	cmd.Run()
	cmd = exec.Command("extra/livego.exe")
	cmd.Run()
}