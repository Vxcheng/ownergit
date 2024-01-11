package main

import (
	_ "ownergit/beego_api/routers"

	"github.com/astaxie/beego"
)

var (
	a = 1
)

// bee run -downdoc=true -gendoc=true
func main() {
	if 1 == 1 || false {

	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
