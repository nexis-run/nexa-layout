package router

import (
	"strings"

	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/assets"
	"nexis.run/nexa-layout/internal/application/http/core"
	userhandler "nexis.run/nexa-layout/internal/application/http/handler/user"
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

	// 接口文档路由
	setDocsRouter(e)

	// 全局中间件
	e.Use(
		rest.CORSMiddleware(
			rest.CORSWithAllowHeaders(
				core.HeaderLayoutUserToken,
			),
		),
		core.ContextMiddleware(),

		// 请求体 / 响应体 dump 日志
		rest.DumpMiddleware(func(c echo.Context) bool {
			// 过滤掉不需要 dump 的路径
			return strings.HasPrefix(c.Request().URL.Path, "/docs") ||
				strings.HasPrefix(c.Request().URL.Path, "/healthz") ||
				strings.HasPrefix(c.Request().URL.Path, "/favicon.ico")
		}),
	)

	// 用户登录
	e.POST("/login", userhandler.Login)

	// 鉴权路由组
	g := e.Group("/user")
	g.Use(
		core.AuthMiddleware(),
	)
	g.GET("/info", userhandler.Info)
}
