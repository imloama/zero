package main

import (
	"fmt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
	"github.com/mazhaoyong/zero/logs"
	"github.com/mazhaoyong/zero/routes"
	"github.com/mazhaoyong/zero/settings"
)

func main() {

	settings.Init()

	logs.Init()

	iris.Favicon("./public/favicon.ico")

	iris.Static("/public", "./public", 1)

	iris.Use(logger.New())

	//crs := cors.New(cors.Options{})

	iris.Use(cors.Default()) // crs

	iris.Use(recovery.Handler)

	routes.Init()

	uri := fmt.Sprintf(":%d", settings.Get().Port)

	logs.Debug("server started on ", uri)

	iris.Listen(uri)
}
