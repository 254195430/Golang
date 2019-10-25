package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (t *HomeController) Get() {
	t.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}
