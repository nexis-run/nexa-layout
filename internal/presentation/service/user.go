package service

import (
	"nexis.run/nexa/kit"

	"nexis.run/nexa-layout/internal/infrastructure/model"
	"nexis.run/nexa-layout/internal/presentation/entity"
	"nexis.run/nexa-layout/internal/presentation/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUser() *UserService {
	return &UserService{
		repo: repository.NewUser(),
	}
}

func (s *UserService) Login(req *entity.UserLoginRequest) (*entity.UserLoginResponse, error) {
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

func (s *UserService) AuthToken(token string) (*model.User, error) {
	if token != "token" {
		return nil, kit.ErrUnauthorized
	}
	return s.repo.FindUserByUsername("demo")
}

func (s *UserService) Info(user *model.User) (*entity.UserInfoResponse, error) {
	return &entity.UserInfoResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
