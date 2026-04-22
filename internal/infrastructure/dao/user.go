package dao

import "nexis.run/nexa-layout/internal/infrastructure/model"

type UserDao struct {
}

func (dao *UserDao) GetUserByUsername(username string) (user *model.User, err error) {
	return &model.User{
		ID:       1,
		Username: username,
		Role:     1,
	}, nil
}

func NewUser() *UserDao {
	return &UserDao{}
}
