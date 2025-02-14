package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/internal/presentation/service"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*LayoutContext)
			if !ok {
				return kit.ErrInvalidContext
			}
			token := c.Request().Header.Get(HeaderLayoutUserToken)
			user, err := service.NewUserServiceImpl().AuthToken(token)
			if err != nil {
				return rest.NewError(http.StatusUnauthorized, err.Error())
			}
			ctx.User = user
			return next(ctx)
		}
	}
}
