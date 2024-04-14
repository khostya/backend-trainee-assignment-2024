package main

import (
	"backend-trainee-assignment-2024/config"
	"backend-trainee-assignment-2024/internal/app"
	"log/slog"
	"os"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	cfg := config.MustConfig()
	app, err := app.NewApp(cfg)
	if err != nil {
		app.Shutdown() // nolint
		slog.Error(err.Error())
		return
	}

	if err := app.Run(); err != nil {
		slog.Error(err.Error())
	}
}
