package bootstrap

import (
	"os"
	"time"

	"go.uber.org/zap"
	"nexis.run/nexa/kit/logger"

	"nexis.run/nexa-layout/internal/config"
	"nexis.run/nexa-layout/internal/server"
)

// Boot 初始化项目
// cfgPath 配置文件路径
func Boot(cfgPath, ver string) {
	// 设置全局时区
	tz := "Asia/Shanghai"
	_ = os.Setenv("TZ", tz)
	loc, _ := time.LoadLocation(tz)
	time.Local = loc

	// 加载配置
	config.Setup(cfgPath, ver)

	// 获取配置
	cfg := config.Get()

	// 初始化日志
	logger.Setup(cfg.Logger)

	zap.L().Info("Boot successed")

	// 启动服务器
	server.Run()
}
