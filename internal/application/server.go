package application

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2"

	"nexis.run/nexa-layout/internal/application/http"
	"nexis.run/nexa-layout/internal/application/micro"
	"nexis.run/nexa-layout/internal/config"
	"nexis.run/nexa-layout/internal/di"
)

// Run 启动双协议服务，等待退出信号并在截止时间内关闭
func Run(ctx context.Context) error {
	cfg := di.C.Config
	app := kratos.New(
		kratos.Context(ctx),
		kratos.Name(cfg.App),
		kratos.Version(config.Version),
		kratos.StopTimeout(30*time.Second),
		kratos.Server(http.Setup(cfg.App, cfg.HTTP.Bind), micro.Setup(cfg.GRPC.Bind)),
	)

	return app.Run()
}
