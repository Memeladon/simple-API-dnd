package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"

	authgrpc "simple-API-dnd/sso/internal/grpc/auth"
)

type App struct {
	logger     *slog.Logger
	GRPCServer *grpc.Server
	port       int
}

func NewApp(logger *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.RegisterServerAPI(gRPCServer)

	return &App{
		logger:     logger,
		GRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	op := "grpcapp.Run"
	log := a.logger.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.GRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() error {
	op := "grpcapp.Stop"

	a.logger.With(slog.String("op", op)).Info("grpc server is stopping")
	a.GRPCServer.GracefulStop()

	return nil
}
