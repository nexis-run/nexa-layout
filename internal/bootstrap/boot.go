package bootstrap

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"nexis.run/nexa/kit/logger"

	"nexis.run/nexa-layout/internal/config"
	"nexis.run/nexa-layout/internal/di"
)

// Boot 初始化项目
// cfgPath 配置文件路径
func Boot(cfgPath string) {
	// 设置全局时区
	tz := "Asia/Shanghai"
	_ = os.Setenv("TZ", tz)
	loc, _ := time.LoadLocation(tz)
	time.Local = loc

	// 加载配置
	config.Setup(cfgPath)

	// 获取配置
	cfg := config.Get()

	// 初始化日志
	logger.Setup(cfg.Logger)

	// 注入依赖
	var err error
	di.C, err = di.Initialize(cfg)
	if err != nil {
		log.Fatalf("依赖注入失败: %v", err)
	}

	// 打印启动信息
	zap.S().Infof("初始化完成, 当前版本号: %s", config.Version)
}
