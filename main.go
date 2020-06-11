package main

import (
	_ "LiveAssembly/routers"
	"os/exec"
	"github.com/astaxie/beego"
)

func init(){
	cmd:= exec.Command("taskkill /f /t /im livego.exe")
	cmd.Run()
	cmd = exec.Command("extra/livego.exe")
	cmd.Run()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
