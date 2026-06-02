package config

import (
	"nexis.run/nexa/kit/configure"
)

// Version 编译期由 -ldflags 注入
var Version string

// Config 项目配置
type Config struct {
	configure.Configure

	BaseURL string

	GRPC struct {
		Bind string
	}

	HTTP struct {
		Bind string
	}
}

// Load 加载配置文件
func Load(p string) (*Config, error) {
	c, err := configure.Load[Config](p)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
