package main

import (
	_ "catapi/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Start the Beego web application
	web.Run()
}
