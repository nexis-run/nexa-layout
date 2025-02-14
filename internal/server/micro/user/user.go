// Copyright (C) micro-layout. 2025-present.
//
// Created at 2025-02-10, by liasica

package user

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"

	layoutv1 "nexis.run/nexa-layout/pb/gen/layout/v1"
)

func Register(s *grpc.Server) {
	layoutv1.RegisterUserServiceServer(s, NewGrpcServer(NewService()))
}
