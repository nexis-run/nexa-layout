package router

import (
	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/assets"
	"nexis.run/nexa-layout/internal/application/http/core"
	"nexis.run/nexa-layout/internal/application/http/handler"
)

// Setup 设置Rest路由
//
// @title				layout API
// @version				1.0
// @servers.url			>>[BASE_URL]<<
// @servers.description	>>[ENV]<< BaseURL
// @doc					https://github.com/swaggo/swag/issues/386 https://github.com/swaggo/swag/issues/548 https://github.com/go-openapi/runtime/blob/master/middleware/redoc.go
func Setup(e *echo.Echo) {
	// 加载模板
	e.Renderer = rest.LoadTemplates(assets.TemplateFS, "templates")

	setDocsRouter(e)

	e.Use(rest.CORSMiddlware(
		rest.CORSWithAllowHeaders(core.HeaderLayoutUserToken),
	))

	// 用户登录
	e.POST("/login", handler.User.Login)

	g := e.Group("/user")
	g.Use(
		core.ContextMiddleware(),
		core.AuthMiddleware(),
	)
	g.GET("/info", handler.User.Info)
}
