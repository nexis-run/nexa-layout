package micro

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"nexis.run/nexa/kit/micro"

	"nexis.run/nexa-layout/internal/application/micro/user"
)

// Setup 初始化微服务
func Setup(address string) *grpc.Server {
	return micro.New(address, func(s *grpc.Server) {
		user.Register(s)
	})
}
