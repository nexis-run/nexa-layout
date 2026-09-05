package bootstrap

import (
	"fmt"
	"time"
	_ "time/tzdata"

	"go.uber.org/zap"
	"nexis.run/nexa/kit/logger"

	"nexis.run/nexa-layout/internal/config"
	"nexis.run/nexa-layout/internal/di"
)

// Boot 初始化项目
// cfgPath 配置文件路径
func Boot(cfgPath string) (err error) {
	// 设置应用时区
	time.Local, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	// 加载配置
	var cfg *config.Config

	cfg, err = config.Load(cfgPath)
	if err != nil {
		err = fmt.Errorf("配置加载失败：%w", err)
		return
	}

	// 初始化日志
	err = logger.Setup(cfg.Logger)
	if err != nil {
		return
	}

	// 注入依赖
	di.C, err = di.Initialize(cfg)
	if err != nil {
		err = fmt.Errorf("依赖注入失败：%w", err)
		return
	}

	// 打印启动信息
	zap.L().Info("应用初始化完成", zap.String("version", config.Version))

	return
}
