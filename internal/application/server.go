// Copyright (C) micro-layout. 2025-present.
//
// Created at 2025-02-10, by liasica

package application

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"nexis.run/nexa/kit/graceful"

	"nexis.run/nexa-layout/internal/application/http"
	"nexis.run/nexa-layout/internal/application/micro"
	"nexis.run/nexa-layout/internal/di"
)

type Server struct {
	graceful.Server

	grpcServer *grpc.Server
	echoServer *echo.Echo
}

func (s *Server) Start() {
	cfg := di.C.Config

	// 启动Rest服务器
	s.echoServer = http.Setup(cfg.App, cfg.HTTP.Bind)

	// 启动Grpc服务器
	s.grpcServer = micro.Setup(cfg.App, cfg.GRPC.Bind)
}

func (s *Server) Stop(ctx context.Context) {
	var wg sync.WaitGroup

	if s.grpcServer != nil {
		wg.Go(func() {
			zap.L().Info("grpc服务器关闭中...")

			s.grpcServer.GracefulStop()
		})
	}

	if s.echoServer != nil {
		wg.Go(func() {
			zap.L().Info("rest服务器关闭中...")

			_ = s.echoServer.Shutdown(ctx)
		})
	}

	wg.Wait()
}

func Run() {
	graceful.Run(&Server{})
}
