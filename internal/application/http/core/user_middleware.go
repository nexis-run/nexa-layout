package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/internal/di"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*Context)
			if !ok {
				return kit.ErrInvalidContext
			}

			token := c.Request().Header.Get(HeaderLayoutUserToken)

			user, err := di.C.Service.User.AuthToken(token)
			if err != nil {
				return rest.NewError(http.StatusUnauthorized, err.Error())
			}

			ctx.User = user

			return next(ctx)
		}
	}
}
