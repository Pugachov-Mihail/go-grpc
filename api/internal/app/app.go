package app

import (
	"log/slog"
	grpcApp "magicMc/api/internal/app/grpc"
	"time"
)

type App struct {
	GRPC *grpcApp.App
}

func New(
	log *slog.Logger,
	portGrpc int,
	storagePath string,
	TokenTTL time.Duration,
) *App {
	grpcApp := grpcApp.New(log, portGrpc)
	return &App{
		GRPC: grpcApp,
	}
}
