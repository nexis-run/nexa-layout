package route

import (
	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/assets"
	"nexis.run/nexa-layout/internal/app/rest/app"
	"nexis.run/nexa-layout/internal/app/rest/controller"
)

// Setup 设置Rest路由
//
//	@title					layout API
//	@version				1.0
//	@servers.url			>>[BASE_URL]<<
//	@servers.description	>>[ENV]<< BaseURL
//	@doc					https://github.com/swaggo/swag/issues/386 https://github.com/swaggo/swag/issues/548 https://github.com/go-openapi/runtime/blob/master/middleware/redoc.go
func Setup(e *echo.Echo) {
	// 加载模板
	e.Renderer = rest.LoadTemplates(assets.TemplateFS, "templates")

	setDocsRoute(e)

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
