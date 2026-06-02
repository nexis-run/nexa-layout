package user

import (
	"github.com/labstack/echo/v4"

	"nexis.run/nexa-layout/internal/application/http/core"
	"nexis.run/nexa-layout/internal/di"
	"nexis.run/nexa-layout/internal/presentation/dto"
)

// Login
//
// @ID		UserLogin
// @Router	/login [POST]
// @Summary	用户登录
// @Tags	User - 用户
// @Produce	json
// @Param	request	body		dto.UserLoginRequest	true	"请求参数"
// @Success	200		{object}	dto.UserLoginResponse	"请求成功"
func Login(c echo.Context) error {
	ctx, req := core.GetContextBindingData[dto.UserLoginRequest](c)

	return ctx.SendResponse(di.C.Service.User.Login(req))
}

// Info
//
// @ID		UserInfo
// @Router	/user/info [GET]
// @Summary	获取用户信息
// @Tags	User - 用户
// @Produce	json
// @Success	200	{object}	dto.UserInfoResponse	"请求成功"
func Info(c echo.Context) error {
	ctx := core.GetContext(c)

	return ctx.SendResponse(di.C.Service.User.Info(ctx.User))
}
