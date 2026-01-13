package config

import (
	"nexis.run/nexa/kit/configure"
)

var (
	instance *Config
	Version  string
)

type Config struct {
	configure.Configure

	BaseUrl string

	Grpc struct {
		Bind string
	}

	Http struct {
		Bind string
	}
}

// Setup 初始化配置
func Setup(p, v string) {
	c, err := configure.Load[Config](p)
	if err != nil {
		panic(err)
	}
	Version = v
	instance = &c
}

// Get 获取配置
func Get() *Config {
	return instance
}
