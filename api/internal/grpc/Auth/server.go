package auth

import (
	"context"
	Auth "magicMc/api/grpc/gen/Auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthApp interface {
	Login(ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)
	RegisterNewUser(ctx context.Context,
		email string,
		password string,
	) (userId int64, err error)
	isAdmin(ctx context.Context, userID int64) (bool, error)
}

type serverAPI struct {
	Auth.UnimplementedAuthApiServer
	authApp AuthApp
}

func Register(gRPC *grpc.Server, auth AuthApp) {
	Auth.RegisterAuthApiServer(gRPC, &serverAPI{authApp: auth})
}

const (
	emptyValue = 0
)

func (s *serverAPI) Login(
	ctx context.Context,
	req *Auth.LoginRequest,
) (*Auth.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	token, err := s.authApp.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		// TODO:
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &Auth.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req Auth.RegisterRequest,
) (*Auth.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}
}

func (s *serverAPI) isAdmin(
	ctx context.Context,
	req Auth.IsAdminRequest,
) *Auth.IsAdminResponse {
	panic("implement me")
}

func validateLogin(req *Auth.LoginRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email empty")
	}

	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "password empty")
	}

	if req.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "app is required")
	}
	return nil
}

func validateRegister(req Auth.RegisterRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email empty")
	}

	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "password empty")
	}
	return nil
}
