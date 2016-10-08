package controllers

//HomeController home controller
type HomeController struct {
	Controller
}

// Index home
//@router / [get]
func (p *HomeController) Index() {
	//TODO
	p.Data["Website"] = "beego.me"
	p.Data["Email"] = "astaxie@gmail.com"
	p.TplName = "index.tpl"
}
