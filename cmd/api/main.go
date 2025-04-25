package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/viniosilva/go-boilerplateapi/config"
	"github.com/viniosilva/go-boilerplateapi/internal/container"
	"github.com/viniosilva/go-boilerplateapi/internal/infrastructure/api"
	"github.com/viniosilva/go-boilerplateapi/internal/infrastructure/db"
)

// @title Ipanema Box API
// @version 1.0
// @description API management for customers and services
// @BasePath /api
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config.LoadConfig: %v", err)
	}

	dbConn, err := db.NewGorm(cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName,
		cfg.DB.User, cfg.DB.Password, cfg.DB.SslMode,
	)
	if err != nil {
		log.Fatalf("db.NewGorm: %v", err)
	}
	defer db.Close(dbConn)

	di := container.New(dbConn)
	server := api.NewServer(di, cfg.App.Host, cfg.App.Port, time.Second*time.Duration(cfg.App.TimeoutSec))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		log.Printf("api listening: %s", server.Addr)

		if err = server.ListenAndServe(); err != nil {
			log.Fatalf("http.ListenAndServe: %v", err)
		}
	}()

	<-quit
	log.Println("server closing")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server.Shutdown: %v", err)
	}

	log.Println("server closed successfully")
}
