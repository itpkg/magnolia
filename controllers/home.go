package controllers

//HomeController home controller
type HomeController struct {
	Controller
}

// Home home
//@router / [get]
func (p *HomeController) Home() {
	//TODO
	p.Data["Title"] = "home"
	p.TplName = "home.tpl"
}
