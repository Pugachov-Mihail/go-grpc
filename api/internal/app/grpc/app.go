package grpcApp

import (
	"fmt"
	"log/slog"
	authGRPC "magicMc/api/internal/grpc/Auth"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log  *slog.Logger
	gRPC *grpc.Server
	port int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gGRPCserver := grpc.NewServer()
	authGRPC.Register(gGRPCserver)

	return &App{
		log:  log,
		gRPC: gGRPCserver,
		port: port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting grpc server", slog.String("addr", l.Addr().String()))

	if err := a.gRPC.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping grpc server", slog.Int("port", a.port))

	a.gRPC.GracefulStop()
}
