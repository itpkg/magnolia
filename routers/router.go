package routers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/itpkg/magnolia/controllers"
	"github.com/itpkg/magnolia/controllers/admin"
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
	beego.Include(&admin.Controller{})
	beego.Include(&ops.Controller{})

	for k, v := range map[string]beego.ControllerInterface{
		"forum":   &forum.Controller{},
		"reading": &reading.Controller{},
		"shop":    &shop.Controller{},
		"opsmail": &ops_mail.Controller{},
		"opsvpn":  &ops_vpn.Controller{},
	} {
		if beego.AppConfig.DefaultBool(fmt.Sprintf("engine%senable", k), true) {
			beego.Include(v)
		}
	}
}
