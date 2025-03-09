package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/a179346/recommendation-system/internal/app/config"
	"github.com/a179346/recommendation-system/internal/app/database/dbhelper"
	"github.com/a179346/recommendation-system/internal/app/server"
	"github.com/a179346/recommendation-system/internal/pkg/console"
	"github.com/a179346/recommendation-system/internal/pkg/graceful"
	"github.com/a179346/recommendation-system/internal/pkg/redishelper"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		console.Errorf("Exit 1: %v", err)
		os.Exit(1)
	}
	console.Info("Exit 0")
}

func run() error {
	db, err := dbhelper.Open()
	if err != nil {
		return fmt.Errorf("opendb.Open error: %w", err)
	}
	defer func() {
		console.Info("Closing db connection...")
		db.Close()
	}()
	db.SetMaxOpenConns(30)
	dbhelper.WaitFor(context.Background(), db)

	redisConfig := config.GetRedisConfig()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	defer func() {
		console.Info("Closing redis connection...")
		redisClient.Close()
	}()
	redishelper.WaitForConnected(context.Background(), redisClient)

	server := server.GetServer(db, redisClient)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()
		console.Info("Shutting down server...")
		if err := server.Shutdown(ctx); err != nil {
			console.Errorf("Error shutting down server: %v", err)
		}
	}()

	serverListenErrCh := make(chan error)
	go func() {
		address := fmt.Sprintf(":%d", config.GetServerConfig().Port)
		if err := server.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverListenErrCh <- fmt.Errorf("Error starting server: %w", err)
		}
	}()

	select {
	case signal := <-graceful.ShutDown():
		console.Infof("Received signal: %v", signal)
		return nil

	case err := <-serverListenErrCh:
		return err
	}
}
