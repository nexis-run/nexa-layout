// Copyright (C) micro-layout. 2025-present.
//
// Created at 2025-02-10, by liasica

package user

import (
	"context"

	"nexis.run/nexa-layout/internal/presentation/dto"
	"nexis.run/nexa-layout/internal/presentation/service"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Login(_ context.Context, req *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	return service.NewUser().Login(req)
}
