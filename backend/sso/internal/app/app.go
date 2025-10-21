package app

import (
	"log/slog"
	grpcapp "simple-API-dnd/sso/internal/app/grpc_app"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	//TODO: инициализировать хранилище
	//TODO: инициализировать auth service

	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
