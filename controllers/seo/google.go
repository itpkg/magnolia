package seo

import "fmt"

//GetGoogle google verify
//@router /google:code([\w]+).html [get]
func (p *Controller) GetGoogle() {
	//TODO select from database
	p.Ctx.WriteString(
		fmt.Sprintf(
			"google-site-verification: google%s.html",
			p.Ctx.Input.Param(":code"),
		),
	)
}
