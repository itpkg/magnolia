package seo

//GetBaidu baidu verify
//@router /baidu_verify_:code([\w]+).html [get]
func (p *Controller) GetBaidu() {
	//TODO select from database
	p.Ctx.WriteString(p.Ctx.Input.Param(":code"))
}
