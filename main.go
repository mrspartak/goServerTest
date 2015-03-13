package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.HttpPort = 8000
	beego.Router("/go", &MainController{})
	beego.Run()
}
