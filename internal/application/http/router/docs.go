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
	"nexis.run/nexa-layout/internal/di"
)

func setDocsRouter(e *echo.Echo) {
	openapiPath := "/docs/openapi.yaml"
	docPath := "/docs"

	e.GET(openapiPath, func(c echo.Context) error {
		envStr := "正式环境"

		u, err := docsBaseURL(c, openapiPath)
		if err != nil {
			return err
		}

		switch di.C.Config.Environment {
		case kit.Staging:
			envStr = "测试环境"
		case kit.Development:
			envStr = "开发环境"
		case kit.Production:
			envStr = "正式环境"
		}

		content := strings.NewReplacer(
			">>[BASE_URL]<<", u.String(), ">>[ENV]<<", envStr,
		).Replace(assets.OpenAPIData)

		return c.Blob(http.StatusOK, "application/yaml; charset=UTF-8", []byte(content))
	})

	e.GET(docPath, func(c echo.Context) error {
		u, err := docsBaseURL(c, docPath)
		if err != nil {
			return err
		}

		specURL := u.JoinPath(openapiPath).String()

		return c.Render(http.StatusOK, "docs.html", map[string]string{
			"title":   di.C.Config.App + " API",
			"specURL": specURL,
		})
	})
}

func docsBaseURL(c echo.Context, endpoint string) (address *url.URL, err error) {
	if configured := di.C.Config.BaseURL; configured != "" {
		address, err = url.Parse(configured)
		return
	}

	address, err = rest.GetRequestURL(c)
	if err != nil {
		return
	}

	address.RawPath = strings.TrimSuffix(address.EscapedPath(), endpoint)
	address.Path = strings.TrimSuffix(address.Path, endpoint)
	address.RawQuery = ""
	address.ForceQuery = false
	address.Fragment = ""
	address.RawFragment = ""

	return
}
