package repository

import "nexis.run/nexa-layout/internal/infrastructure/model"

type UserRepository struct {
}

func (u *UserRepository) FindUserByUsername(username string) (user *model.User, err error) {
	return &model.User{
		ID:       1,
		Username: username,
		Role:     1,
	}, nil
}

func NewUser() *UserRepository {
	return &UserRepository{}
}
