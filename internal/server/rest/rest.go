package rest

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"orba.plus/nexa/kit/rest"

	"orba.plus/nexa-layout/internal/server/rest/route"
)

// Setup 初始化Rest服务
func Setup(name, address string) (echoServer *echo.Echo) {
	var ch chan error
	echoServer, ch = rest.Run(name, address, func(e *echo.Echo) {
		route.Setup(e)
	})
	go func() {
		if err := <-ch; err != nil {
			zap.L().Fatal("rest服务启动失败", zap.Error(err))
		}
	}()
	return
}
