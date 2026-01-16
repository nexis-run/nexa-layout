package app

import (
	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/internal/infrastructure/model"
)

type Context struct {
	*rest.Context

	User *model.User
}

func NewContext(c *rest.Context) *Context {
	return &Context{Context: c}
}

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(NewContext(c.(*rest.Context)))
		}
	}
}

func GetContext(c echo.Context) *Context {
	switch v := c.(type) {
	case *Context:
		return v
	default:
		return NewContext(rest.GetContext(c))
	}
}

func ContextBinding[T any](c echo.Context) (*Context, *T) {
	ctx := GetContext(c)
	req := new(T)
	ctx.BindValidate(req)
	return ctx, req
}
