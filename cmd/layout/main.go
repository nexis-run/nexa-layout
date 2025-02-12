package main

import (
	"flag"

	"orba.plus/nexa-layout/internal/bootstrap"
)

var Version = "v1.0.0"

func main() {
	var cfg string
	flag.StringVar(&cfg, "config", "config/config.yaml", "配置文件")
	flag.Parse()

	bootstrap.Boot(cfg, Version)
}
