package service

import (
	"nexis.run/nexa-layout/internal/domain/entity"
	"nexis.run/nexa-layout/internal/infrastructure/model"
)

// UserService 抽象用户服务
type UserService interface {
	// Login 登录
	Login(entity *entity.UserLoginRequest) (*entity.UserLoginResponse, error)

	// AuthToken 验证token
	AuthToken(token string) (*model.User, error)

	// Info 获取用户信息
	Info(user *model.User) (*entity.UserInfoResponse, error)
}
