package main

import (
	"github.com/astaxie/beego"
	_ "github.com/victorfengming/gochatroom/routers"
)

func main() {
	beego.Run()
}
