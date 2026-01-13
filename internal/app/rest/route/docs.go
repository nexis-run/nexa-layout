// Copyright (C) moneta. 2025-present.
//
// Created at 2025-08-04, by liasica

package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit"

	"nexis.run/nexa-layout/assets"
	"nexis.run/nexa-layout/internal/config"
)

func setDocsRoute(e *echo.Echo) {
	docPath := "/docs/openapi.yaml"

	e.GET(docPath, func(c echo.Context) error {
		envStr := "正式环境"
		baseUrl := config.Get().BaseUrl
		if baseUrl == "" {
			baseUrl = fmt.Sprintf("http://%s", c.Request().Host)
		}

		switch config.Get().Environment {
		case kit.Staging:
			envStr = "测试环境"
		case kit.Development:
			envStr = "开发环境"
		case kit.Production:
			envStr = "正式环境"
		}

		assets.OpenApiData = strings.ReplaceAll(assets.OpenApiData, ">>[BASE_URL]<<", baseUrl)
		assets.OpenApiData = strings.ReplaceAll(assets.OpenApiData, ">>[ENV]<<", envStr)

		return c.String(http.StatusOK, assets.OpenApiData)
	})

	e.GET("/docs", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs.html", map[string]string{
			"title":   fmt.Sprintf("「%s」文档", config.Get().App),
			"specURL": docPath,
		})
	})
}
