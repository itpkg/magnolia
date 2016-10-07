package routers

import (
	"github.com/astaxie/beego"
	"github.com/itpkg/magnolia/controllers"
	"github.com/itpkg/magnolia/controllers/auth"
	"github.com/itpkg/magnolia/controllers/forum"
	"github.com/itpkg/magnolia/controllers/ops"
	ops_mail "github.com/itpkg/magnolia/controllers/ops/mail"
	ops_vpn "github.com/itpkg/magnolia/controllers/ops/vpn"
	"github.com/itpkg/magnolia/controllers/reading"
	"github.com/itpkg/magnolia/controllers/seo"
	"github.com/itpkg/magnolia/controllers/shop"
)

func init() {
	beego.Include(&controllers.HomeController{})
	beego.Include(&seo.Controller{})
	beego.Include(&auth.Controller{})
	beego.Include(&forum.Controller{})
	beego.Include(&reading.Controller{})
	beego.Include(&shop.Controller{})
	beego.Include(&ops.Controller{})
	beego.Include(&ops_mail.Controller{})
	beego.Include(&ops_vpn.Controller{})
}
