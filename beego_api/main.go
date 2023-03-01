package main

import (
	_ "ownergit/beego_api/routers"

	"github.com/astaxie/beego"
)

// bee run -downdoc=true -gendoc=true
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
