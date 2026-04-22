// Copyright (C) moneta. 2025-present.
//
// Created at 2025-08-04, by liasica

package router

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/assets"
	"nexis.run/nexa-layout/internal/config"
)

func setDocsRouter(e *echo.Echo) {
	openapiPath := "/docs/openapi.yaml"
	docPath := "/docs"

	e.GET(openapiPath, func(c echo.Context) error {
		envStr := "正式环境"

		u, err := rest.GetRequestUrl(c)
		if err != nil {
			return err
		}

		baseUrl := strings.Replace(u.String(), openapiPath, "", 1)

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

	e.GET(docPath, func(c echo.Context) error {
		u, err := rest.GetRequestUrl(c)
		if err != nil {
			return err
		}

		baseUrl := strings.Replace(u.String(), docPath, "", 1)
		specURL, _ := url.JoinPath(baseUrl, openapiPath)

		return c.Render(http.StatusOK, "docs.html", map[string]string{
			"title":   "订单管理系统 API",
			"specURL": specURL,
		})
	})
}
