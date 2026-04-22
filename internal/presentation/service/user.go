package service

import (
	"nexis.run/nexa/kit"

	"nexis.run/nexa-layout/internal/infrastructure/dao"
	"nexis.run/nexa-layout/internal/infrastructure/model"
	"nexis.run/nexa-layout/internal/presentation/dto"
)

type UserService struct {
	repo *dao.UserDao
}

func NewUser() *UserService {
	return &UserService{
		repo: dao.NewUser(),
	}
}

func (s *UserService) Login(req *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	user, err := s.repo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	res := &dto.UserLoginResponse{
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
	return s.repo.GetUserByUsername("demo")
}

func (s *UserService) Info(user *model.User) (*dto.UserInfoResponse, error) {
	return &dto.UserInfoResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
