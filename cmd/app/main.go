package main

import (
	"cat_service/internal/app"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

// @title Cat Service API
// @version 1.0
// @description API for managing cats and feedback
// @host localhost:8080
// @BasePath /
func main() {
	fx.New(
		app.ModuleDB(),
		app.ModuleRepositories(),
		app.ModuleServices(),
		app.ModuleHandlers(),
		app.ModuleSwagger(),
		app.ModuleApp(),
	).Run()
}
