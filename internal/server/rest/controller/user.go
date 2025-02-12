package controller

import (
	"github.com/labstack/echo/v4"
	"orba.plus/nexa/kit/rest"

	"orba.plus/nexa-layout/internal/domain/entity"
	"orba.plus/nexa-layout/internal/presentation/service"
	"orba.plus/nexa-layout/internal/server/rest/app"
)

type user struct {
}

var User = new(user)

func (*user) Login(c echo.Context) (err error) {
	ctx, req := rest.BaseContextBinding[entity.UserLoginRequest](c)
	return ctx.SendResponse(service.NewUserServiceImpl().Login(req))
}

func (*user) Info(c echo.Context) (err error) {
	ctx := app.Context(c)
	return ctx.SendResponse(service.NewUserServiceImpl().Info(ctx.User))
}
