package routers

import (
	"github.com/astaxie/beego"
	"github.com/itpkg/magnolia/controllers"
)

func init() {
	beego.Include(&controllers.HomeController{})
}
