package service

import (
	"nexis.run/nexa/kit"

	"nexis.run/nexa-layout/internal/domain/entity"
	"nexis.run/nexa-layout/internal/infrastructure/model"
	"nexis.run/nexa-layout/internal/presentation/repository"
)

type UserServiceImpl struct {
	repo *repository.UserRepositoryImpl
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		repo: repository.NewUserRepository(),
	}
}

func (s *UserServiceImpl) Login(req *entity.UserLoginRequest) (*entity.UserLoginResponse, error) {
	user, err := s.repo.FindUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	res := &entity.UserLoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Token:    "token",
	}
	return res, nil
}

func (s *UserServiceImpl) AuthToken(token string) (*model.User, error) {
	if token != "token" {
		return nil, kit.ErrUnauthorized
	}
	return s.repo.FindUserByUsername("demo")
}

func (s *UserServiceImpl) Info(user *model.User) (*entity.UserInfoResponse, error) {
	return &entity.UserInfoResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
