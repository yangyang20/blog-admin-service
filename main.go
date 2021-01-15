package main

import (
	"blog-admin/app"
	_ "blog-admin/boot"
	_ "blog-admin/router"
)

func main() {
	//g.Server().Run()
	app.Run()
}
