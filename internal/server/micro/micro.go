package micro

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.uber.org/zap"
	"nexis.run/nexa/kit/micro"

	"nexis.run/nexa-layout/internal/server/micro/user"
)

// Setup 初始化微服务
func Setup(name, address string) (server *grpc.Server) {
	var ch chan error
	server, ch = micro.Run(name, address, func(s *grpc.Server) {
		user.Register(s)
	})
	go func() {
		if err := <-ch; err != nil {
			zap.L().Fatal(err.Error())
		}
	}()
	return
}
