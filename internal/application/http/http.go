package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/internal/application/http/router"
)

// Server 实现应用的 HTTP 启动与关闭接口
type Server struct {
	echoServer *echo.Echo
	address    string
}

func Setup(app, address string) *Server {
	echoServer := rest.New(app, router.Setup)
	echoServer.Server.ReadHeaderTimeout = 5 * time.Second
	echoServer.Server.IdleTimeout = time.Minute

	return &Server{echoServer: echoServer, address: address}
}

func (server *Server) Start(_ context.Context) error {
	err := server.echoServer.Start(server.address)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (server *Server) Stop(ctx context.Context) error {
	err := server.echoServer.Shutdown(ctx)
	if err != nil {
		return errors.Join(err, server.echoServer.Close())
	}

	return nil
}
