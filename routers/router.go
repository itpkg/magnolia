package routers

import (
	"github.com/itpkg/magnolia/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
