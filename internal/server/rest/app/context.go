package app

import (
	"github.com/labstack/echo/v4"
	"orba.plus/nexa/kit/rest"

	"orba.plus/nexa-layout/internal/infrastructure/model"
)

type LayoutContext struct {
	*rest.Context

	User *model.User
}

func NewLayoutContext(c *rest.Context) *LayoutContext {
	return &LayoutContext{Context: c}
}

func LayoutContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(NewLayoutContext(c.(*rest.Context)))
		}
	}
}

func Context(c echo.Context) *LayoutContext {
	switch v := c.(type) {
	case *LayoutContext:
		return v
	default:
		return NewLayoutContext(rest.GetContext(c))
	}
}

func ContextBinding[T any](c echo.Context) (*LayoutContext, *T) {
	ctx := Context(c)
	req := new(T)
	ctx.BindValidate(req)
	return ctx, req
}
