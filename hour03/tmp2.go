package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (a *HomeController) Get() {
	a.Ctx.WriteString("hello world")
}
func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}
