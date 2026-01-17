package bootstrap

import (
	"os"
	"time"

	"go.uber.org/zap"
	"nexis.run/nexa/kit/logger"

	"nexis.run/nexa-layout/internal/config"
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

	// 打印启动信息
	zap.S().Infof("加载完成, 当前版本号: %s", config.Version)
}
