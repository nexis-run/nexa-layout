package route

import (
	"github.com/labstack/echo/v4"
	"orba.plus/nexa/kit/rest"

	"orba.plus/nexa-layout/internal/server/rest/app"
	"orba.plus/nexa-layout/internal/server/rest/controller"
)

func Setup(e *echo.Echo) {
	e.Use(rest.CORSMiddlware(
		rest.CORSWithAllowHeaders(app.HeaderLayoutUserToken),
	))

	// 用户登录
	e.POST("/login", controller.User.Login)

	g := e.Group("/user")
	g.Use(
		app.LayoutContextMiddleware(),
		app.AuthMiddleware(),
	)
	g.GET("/info", controller.User.Info)
}
