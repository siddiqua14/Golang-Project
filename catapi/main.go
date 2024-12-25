package main

import (
	_ "catapi/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
