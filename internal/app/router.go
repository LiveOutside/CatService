package app

import (
	"cat_service/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	FeedbackHandler *handlers.FeedbackHandler
	HomeHandler     *handlers.HomeHandler
	CatHandler      *handlers.CatHandler
}

func RegisterRoutes(app *fiber.App, handlers Handlers) {
	app.Get("/", handlers.HomeHandler.Get)

	app.Get("/api/cats", handlers.CatHandler.Get)
	app.Get("/api/cats/feedback", handlers.FeedbackHandler.Get)
	app.Get("/api/cats/:id", handlers.CatHandler.GetByID)
	app.Get("/views/cats/internal/:id", handlers.CatHandler.GetViewByID)
	app.Post("/api/cats", handlers.CatHandler.Post)
}
