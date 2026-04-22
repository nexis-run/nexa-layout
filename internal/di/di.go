package di

import (
	"github.com/google/wire"
	"nexis.run/nexa-layout/internal/config"
	"nexis.run/nexa-layout/internal/infrastructure/dao"
)

type Container struct {
	// 配置中心
	Config *config.Config

	// 集成服务
	Integration *Intergration

	// Dao
	Dao *Dao
}

// Intergration 集成服务
// 集成所有第三方或外部服务
type Intergration struct {
}

type Dao struct {
	User *dao.UserDao
}

var daoProviderSet = wire.NewSet(
	dao.NewUser,

	wire.Struct(new(Dao), "*"),
)
