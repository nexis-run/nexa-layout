package handler

import (
	"github.com/labstack/echo/v4"

	"nexis.run/nexa-layout/internal/application/http/core"
	"nexis.run/nexa-layout/internal/presentation/dto"
	"nexis.run/nexa-layout/internal/presentation/service"
)

type UserHandler struct {
}

var User = new(UserHandler)

// Login
//
// @ID		UserLogin
// @Router	/login [POST]
// @Summary	用户登录
// @Tags	User - 用户
// @Produce	json
// @Param	request	body		dto.UserLoginRequest	true	"请求参数"
// @Success	200		{object}	dto.UserLoginResponse	"请求成功"
func (*UserHandler) Login(c echo.Context) (err error) {
	ctx, req := core.GetContextBindingData[dto.UserLoginRequest](c)
	return ctx.SendResponse(service.NewUser().Login(req))
}

// Info
//
// @ID		UserInfo
// @Router	/UserHandler/info [GET]
// @Summary	获取用户信息
// @Tags	User - 用户
// @Produce	json
// @Success	200	{object}	dto.UserInfoResponse	"请求成功"
func (*UserHandler) Info(c echo.Context) (err error) {
	ctx := core.GetContext(c)
	return ctx.SendResponse(service.NewUser().Info(ctx.User))
}
