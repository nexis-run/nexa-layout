// Copyright (C) micro-layout. 2025-present.
//
// Created at 2025-02-10, by liasica

package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"nexis.run/nexa/kit/graceful"

	"nexis.run/nexa-layout/internal/app/micro"
	"nexis.run/nexa-layout/internal/app/rest"
	"nexis.run/nexa-layout/internal/config"
)

type Server struct {
	graceful.Server

	grpcServer *grpc.Server
	echoServer *echo.Echo
}

func (s *Server) Start() {
	cfg := config.Get()

	// 启动Rest服务器
	s.echoServer = rest.Setup(cfg.App, cfg.Http.Bind)

	// 启动Grpc服务器
	s.grpcServer = micro.Setup(cfg.App, cfg.Grpc.Bind)
}

func (s *Server) Stop(ctx context.Context) {
	if s.grpcServer != nil {
		zap.L().Info("grpc服务器关闭中...")
		s.grpcServer.GracefulStop()
	}
	if s.echoServer != nil {
		zap.L().Info("rest服务器关闭中...")
		_ = s.echoServer.Shutdown(ctx)
	}
}

func Run() {
	graceful.Run(&Server{})
}
