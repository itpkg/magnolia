package controllers

import "github.com/astaxie/beego"

//HomeController home controller
type HomeController struct {
	beego.Controller
}

// Index home
//@router / [get]
func (p *HomeController) Index() {
	//TODO
	p.Data["Website"] = "beego.me"
	p.Data["Email"] = "astaxie@gmail.com"
	p.TplName = "index.tpl"
}
