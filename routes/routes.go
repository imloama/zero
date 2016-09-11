// 路由总配置
package routes

import (
	"github.com/kataras/iris"
	//"github.com/mazhaoyong/zero/settings"
	"github.com/mazhaoyong/zero/utils"
)

func Init() {
	myJwtMiddleware := utils.JWTMiddleWare()

	iris.Get("/api/*", myJwtMiddleware.Serve, func(ctx *iris.Context) {

	})

	iris.Get("/", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{"name": "iris"})
	})

	iris.Get("/login", func(ctx *iris.Context) {
		ctx.Render("login.html", iris.Map{"Title": "Login Page"})
	})

	iris.Post("/login", func(ctx *iris.Context) {
		secret := ctx.PostValue("secret")
		ctx.Session().Set("secret", secret)

		ctx.Redirect("/user")
	})

}
