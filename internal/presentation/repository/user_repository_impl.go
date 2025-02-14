package repository

import "nexis.run/nexa-layout/internal/infrastructure/model"

type UserRepositoryImpl struct {
}

func (u UserRepositoryImpl) FindUserByUsername(username string) (user *model.User, err error) {
	return &model.User{
		ID:       1,
		Username: username,
		Role:     1,
	}, nil
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}
