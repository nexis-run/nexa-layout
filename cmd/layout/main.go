// Copyright (C) oos. 2026-present.
//
// Created at 2026-01-09, by liasica

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"nexis.run/nexa/kit/logger"

	"nexis.run/nexa-layout/cmd/layout/internal"
	"nexis.run/nexa-layout/internal/bootstrap"
	"nexis.run/nexa-layout/internal/config"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() (err error) {
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = errors.Join(err, logger.Close(ctx))
	}()

	var cfg string

	cmd := cobra.Command{
		Use:               "layout",
		Short:             "nexis.run layout",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Version:           config.Version,
		SilenceUsage:      true,
		SilenceErrors:     true,
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			return bootstrap.Boot(cfg)
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

	err = cmd.Execute()

	return
}
