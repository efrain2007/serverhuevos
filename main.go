package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	routes.router(app)

}
