package micro

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.uber.org/zap"
	"nexis.run/nexa/kit/micro"

	"nexis.run/nexa-layout/internal/application/micro/user"
)

// Setup 初始化微服务
func Setup(app, address string) *grpc.Server {
	server, ch := micro.Run(app, address, func(s *grpc.Server) {
		user.Register(s)
	})

	go func() {
		if err := <-ch; err != nil {
			zap.L().Fatal("grpc服务启动失败", zap.Error(err))
		}
	}()

	return server
}
