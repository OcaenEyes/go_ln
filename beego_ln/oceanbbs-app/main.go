package main

import (
	_ "oceanbbs-app/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

