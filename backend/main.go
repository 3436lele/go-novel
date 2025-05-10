package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	interfaces "github.com/3436lele/go-novel/backend/command/interface"
	"github.com/3436lele/go-novel/backend/config"
	_ "github.com/kjk/common/filerotate"
)

var (
	programLogLevel = new(slog.LevelVar) // Info by default
)

func main() {
	log := slog.With(slog.String("domain", "main"))

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		log := log.With(slog.String("goroutine", "stop_signal"))

		quit := make(chan os.Signal, 1)
		defer close(quit)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		signal := <-quit

		log.Info("received signal", slog.String("signal", signal.String()))
	}()

	defer func() {
		if err := recover(); err != nil {
			log.Error("panic", "error", err)
		}
	}()

	run(ctx)
}

func run(ctx context.Context) {
	log := slog.With(slog.String("domain", "main"))

	env := os.Getenv("ENV")
	var configPath string
	if env == "" {
		configPath = "config/config.dev.yaml"
	} else {
		configPath = "config/config." + env + ".yaml"
	}

	cfg, err := config.InitConfig(configPath)
	if err != nil {
		panic("failed to initialize config: " + err.Error())
	}

	log.Info("Initializing server ...")
	server := interfaces.NewServer(cfg)

	if err := server.Start(); err != nil {
		log.Error("server error", "error", err)
	}
}
