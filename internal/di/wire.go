// Copyright (C) nexa-layout. 2026-present.
//
// Created at 2026-01-25, by liasica

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"nexis.run/nexa-layout/internal/config"
)

// Initialize 初始化依赖注入容器
func Initialize(cfg *config.Config) (*Container, error) {
	wire.Build(
		// 集成服务
		integrationProviderSet,

		// DAO
		daoProviderSet,

		// Container
		wire.Struct(new(Container), "*"),
	)

	return &Container{}, nil
}
