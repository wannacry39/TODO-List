package main

import (
	"TODO_App/internal/config"
	"TODO_App/internal/storage/sqlite"
	"log/slog"
	"os"
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

	// o1 := todo.NewTODO("PArty", "2024-08-14", "17:00:00")
	// res, err := storage.AddTODO(o1)
	// if err != nil {
	// 	log.Error("Error during adding event")
	// 	os.Exit(1)
	// }

	log.Info("Event added")
	storage.GetTodayTODOS()
	// fmt.Println(res)

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
