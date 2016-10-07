package utils

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

//Controller controller
type Controller interface {
	Abort(code string)
}

//InternalServerError check error
func InternalServerError(c Controller, e error) {
	if e != nil {
		beego.Error(e)
		c.Abort(strconv.Itoa(http.StatusInternalServerError))
	}
}
