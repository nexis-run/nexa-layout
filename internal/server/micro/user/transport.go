// Copyright (C) micro-layout. 2025-present.
//
// Created at 2025-02-10, by liasica

package user

import (
	"context"

	"orba.plus/nexa-layout/internal/domain/entity"
	layoutv1 "orba.plus/nexa-layout/pb/gen/layout/v1"
)

type GrpcServer struct {
	layoutv1.UnimplementedUserServiceServer

	svc *Service
}

func (g *GrpcServer) Login(ctx context.Context, request *layoutv1.LoginRequest) (*layoutv1.LoginResponse, error) {
	resp, err := g.svc.Login(ctx, &entity.UserLoginRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return &layoutv1.LoginResponse{
		Id:       resp.ID,
		Token:    resp.Token,
		Username: resp.Username,
		Role:     uint32(resp.Role),
	}, nil
}

func NewGrpcServer(svc *Service) *GrpcServer {
	return &GrpcServer{
		svc: svc,
	}
}
