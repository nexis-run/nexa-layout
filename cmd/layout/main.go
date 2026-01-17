// Copyright (C) oos. 2026-present.
//
// Created at 2026-01-09, by liasica

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"nexis.run/nexa-layout/cmd/layout/internal"
	"nexis.run/nexa-layout/internal/bootstrap"
	"nexis.run/nexa-layout/internal/config"
)

func main() {
	var cfg string

	cmd := cobra.Command{
		Use:               "layout",
		Short:             "nexis.run layout",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Version:           config.Version,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			bootstrap.Boot(cfg)
		},
	}

	appGroup, appCommand := internal.App()

	cmd.AddGroup(
		appGroup,
	)

	cmd.AddCommand(
		appCommand,
	)

	cmd.PersistentFlags().StringVarP(&cfg, "config", "c", "config/config.yaml", "配置文件")

	err := cmd.Execute()
	if err != nil {
		fmt.Printf("command execution failed: %v\n", err)
		os.Exit(1)
	}
}
