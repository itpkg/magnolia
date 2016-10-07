package seo

import "fmt"

//GetRobots robots.txt
//@router /robots.txt [get]
func (p *Controller) GetRobots() {
	//TODO select from database
	home := "https://www.change-me.com"
	txt := `
  %s
  `
	p.Ctx.WriteString(
		fmt.Sprintf(
			txt,
			home,
		),
	)
}
