// Copyright (C) oos. 2026-present.
//
// Created at 2026-01-17, by liasica

package internal

import (
	"github.com/spf13/cobra"

	"nexis.run/nexa-layout/internal/application"
)

func App() (*cobra.Group, *cobra.Command) {
	g := &cobra.Group{
		ID:    "app",
		Title: "服务端命令",
	}

	cmd := &cobra.Command{
		Use:               "app",
		Short:             "启动服务端",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		GroupID:           g.ID,
		Args:              cobra.NoArgs,
		RunE: func(command *cobra.Command, _ []string) error {
			return application.Run(command.Context())
		},
	}

	return g, cmd
}
