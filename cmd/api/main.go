package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/viniosilva/go-boilerplateapi/config"
	"github.com/viniosilva/go-boilerplateapi/internal/container"
	"github.com/viniosilva/go-boilerplateapi/internal/infrastructure/api"
	"github.com/viniosilva/go-boilerplateapi/internal/infrastructure/db"
	"github.com/viniosilva/go-boilerplateapi/pkg/otel"
)

// @title Ipanema Box API
// @version 1.0
// @description API management for customers and services
// @BasePath /api
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config.LoadConfig: %v", err)
	}

	otelShutdown, err := otel.SetupOTelSDK(ctx, cfg.App.Name, cfg.Otel.Traces.Endpoint, cfg.Otel.Metrics.Endpoint)
	if err != nil {
		log.Fatalf("otel.SetupOTelSDK: %v", err)
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	dbConn, err := db.NewGorm(cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName,
		cfg.DB.User, cfg.DB.Password, cfg.DB.SslMode,
	)
	if err != nil {
		log.Fatalf("db.NewGorm: %v", err)
	}
	defer db.Close(dbConn)

	di := container.New(dbConn)
	srv := api.NewServer(di, cfg.App.Name, cfg.App.Host, cfg.App.Port, cfg.Swagger.Addr, time.Second*time.Duration(cfg.App.TimeoutSec))

	srvErr := make(chan error, 1)

	go func() {
		slog.Info("api listening", slog.String("address", srv.Addr))
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case err = <-srvErr:
		log.Fatalf("srv.ListenAndServe: %v", err)
	case <-ctx.Done():
		stop()
	}

	err = srv.Shutdown(context.Background())

	log.Println("server closed successfully")
}
