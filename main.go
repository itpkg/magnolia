package main

import (
	"github.com/astaxie/beego"
	_ "github.com/itpkg/magnolia/routers"
	_ "github.com/lib/pq"
)

func main() {
	beego.Run()
}
