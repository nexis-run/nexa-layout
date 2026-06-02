// Copyright (C) nexa-layout. 2026-present.
//
// Created at 2026-01-24, by liasica

package di

import (
	"github.com/google/wire"

	"nexis.run/nexa-layout/internal/config"
	"nexis.run/nexa-layout/internal/infrastructure/dao"
)

// C 全局依赖容器, 由 Initialize 在 bootstrap 阶段初始化
var C *Container

// Container 依赖容器
type Container struct {
	// 配置中心
	Config *config.Config

	// 集成服务
	Integration *Integration

	// Dao
	Dao *Dao
}

// Integration 集成服务
// 集成所有第三方或外部服务
// 示例：阿里云 OSS、短信、验证码、Kafka 等
type Integration struct {
}

// integrationProviderSet 集成服务 Provider 集合
// 后续新增集成时, 将对应的 provider 追加到此处
var integrationProviderSet = wire.NewSet(
	wire.Struct(new(Integration), "*"),
)

// Dao 数据访问对象聚合
type Dao struct {
	User *dao.UserDao
}

// daoProviderSet DAO Provider 集合
// 后续新增 DAO 时, 将对应的构造函数追加到此处
var daoProviderSet = wire.NewSet(
	dao.NewUser,

	wire.Struct(new(Dao), "*"),
)
