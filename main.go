package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/itpkg/magnolia/routers"
	_ "github.com/lib/pq"
)

func main() {
	beego.Run()
}

func init() {
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("dburl"))
}
