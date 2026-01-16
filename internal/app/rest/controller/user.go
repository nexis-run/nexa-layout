package controller

import (
	"github.com/labstack/echo/v4"
	"nexis.run/nexa/kit/rest"

	"nexis.run/nexa-layout/internal/app/rest/app"
	"nexis.run/nexa-layout/internal/presentation/entity"
	"nexis.run/nexa-layout/internal/presentation/service"
)

type user struct {
}

var User = new(user)

// Login
//
//	@ID			UserLogin
//	@Router		/login [POST]
//	@Summary	用户登录
//	@Tags		User - 用户
//	@Produce	json
//	@Param		request	body		entity.UserLoginRequest		true	"请求参数"
//	@Success	200		{object}	entity.UserLoginResponse	"请求成功"
func (*user) Login(c echo.Context) (err error) {
	ctx, req := rest.ContextBinding[entity.UserLoginRequest](c)
	return ctx.SendResponse(service.NewUser().Login(req))
}

// Info
//
//	@ID			UserInfo
//	@Router		/user/info [GET]
//	@Summary	获取用户信息
//	@Tags		User - 用户
//	@Produce	json
//	@Success	200	{object}	entity.UserInfoResponse	"请求成功"
func (*user) Info(c echo.Context) (err error) {
	ctx := app.GetContext(c)
	return ctx.SendResponse(service.NewUser().Info(ctx.User))
}
