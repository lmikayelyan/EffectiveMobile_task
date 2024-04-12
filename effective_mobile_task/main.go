package main

import (
	"effective_mobile_task/internal/app"
	"effective_mobile_task/internal/config"
)

// @title Cars-Catalog API
// @version 1.0
// @description Cars catalog API for Effective Mobile test task.

// @host localhost:8000
// @BasePath /

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}
