package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/internal/application/http/router"
)

// Setup 初始化Rest服务
func Setup(app, address string) *echo.Echo {
	echoServer, ch := rest.Run(app, address, func(e *echo.Echo) {
		router.Setup(e)
	})

	go func() {
		if err := <-ch; err != nil {
			zap.L().Fatal("rest服务启动失败", zap.Error(err))
		}
	}()

	return echoServer
}
