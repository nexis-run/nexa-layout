// Copyright (C) micro-layout. 2025-present.
//
// Created at 2025-02-10, by liasica

package user

import (
	"context"

	"orba.plus/nexa-layout/internal/domain/entity"
	"orba.plus/nexa-layout/internal/presentation/service"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Login(_ context.Context, req *entity.UserLoginRequest) (*entity.UserLoginResponse, error) {
	return service.NewUserServiceImpl().Login(req)
}
