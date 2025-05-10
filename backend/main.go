package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/kjk/common/filerotate"
)

var (
	programLogLevel = new(slog.LevelVar) // Info by default
)

func main() {
	log := slog.With(slog.String("program", "main"))

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

	select {
	case <-ctx.Done():
		log.Info("context done")
	}

}
