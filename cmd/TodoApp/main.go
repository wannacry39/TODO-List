package main

import (
	"TODO_App/internal/config"
	add "TODO_App/internal/http-server/handlers/TODOS/Add"
	gettoday "TODO_App/internal/http-server/handlers/TODOS/GetToday"
	"TODO_App/internal/http-server/middleware/logger"
	"TODO_App/internal/storage/sqlite"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	EnvLocal = "local"
	DevEnv   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := SetupLogger(cfg.Env)
	log.Info("Starting App...", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled.")

	storage, err := sqlite.New(cfg.Storagepath)
	if err != nil {
		log.Error("Failed to create database")
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID) //middleware
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)

	router.Post("/Add", add.New(log, storage))
	router.Get("/Today", gettoday.Get(log, storage))

	log.Info("Starting server", slog.String("Address: ", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.Idletimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case DevEnv:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}
