package main

import (
	"backend-trainee-assignment-2024/config"
	"backend-trainee-assignment-2024/internal/app"
	"log"
)

func main() {
	cfg := config.MustConfig()
	app, err := app.NewApp(cfg)
	if err != nil {
		app.Shutdown() // nolint
		log.Fatal(err)
	}
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}
