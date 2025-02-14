package repository

import (
	"nexis.run/nexa-layout/internal/infrastructure/model"
)

// UserRepository 抽象用户仓库
type UserRepository interface {
	// FindUserByUsername 根据用户名查找用户
	FindUserByUsername(username string) (user *model.User, err error)
}
