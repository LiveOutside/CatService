package app

import (
	catdocs "cat_service/docs"
	"cat_service/internal/handlers"
	gencats "cat_service/internal/repositories/gen/cats"
	genfeedback "cat_service/internal/repositories/gen/feedback"
	"cat_service/internal/services/cats"
	"cat_service/internal/services/feedback"
	"cat_service/internal/services/home"
	"cat_service/pkg/db/postgresql"
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	fiberswagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/fx"
)

func ModuleDB() fx.Option {
	return fx.Provide(
		func() (*pgxpool.Pool, error) {
			return postgresql.InitDB(context.Background(), os.Getenv("POSTGRES_DSN"))
		},
	)
}

func ModuleRepositories() fx.Option {
	return fx.Provide(
		fx.Annotate(gencats.New, fx.As(new(gencats.Querier))),
		fx.Annotate(genfeedback.New, fx.As(new(genfeedback.Querier))),
	)
}

func ModuleServices() fx.Option {
	return fx.Provide(
		func(db *pgxpool.Pool, queries genfeedback.Querier) *feedback.Service {
			return feedback.NewService(db, queries, 5*time.Second)
		},
		func(db *pgxpool.Pool, queries gencats.Querier) *cats.Service {
			return cats.NewService(db, queries, 5*time.Second)
		},
		func() *home.Service {
			return home.NewService("https://api.thecatapi.com/v1/images/search")
		},
	)
}

func ModuleSwagger() fx.Option {
	return fx.Options(
		fx.Invoke(
			func(app *fiber.App) {
				host := os.Getenv("SWAGGER_HOST")
				if host == "" {
					host = "localhost:8080"
				}
				catdocs.SwaggerInfo.Host = host
				app.Get("/swagger/*", fiberswagger.WrapHandler)
			},
		),
	)
}

func ModuleHandlers() fx.Option {
	return fx.Provide(
		handlers.NewFeedbackHandler,
		handlers.NewHomeHandler,
		handlers.NewCatHandler,
	)
}

func ModuleApp() fx.Option {
	return fx.Options(
		fx.Provide(NewApp),
		fx.Invoke(
			func(
				lc fx.Lifecycle,
				application *fiber.App,
				catHandler *handlers.CatHandler,
				feedbackHandler *handlers.FeedbackHandler,
				homeHandler *handlers.HomeHandler,
			) {
				RegisterRoutes(application, Handlers{
					CatHandler:      catHandler,
					FeedbackHandler: feedbackHandler,
					HomeHandler:     homeHandler,
				})

				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							if err := application.Listen(":8080"); err != nil {
								log.Errorf("Failed to start listening (fiber): %w", err)
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return application.ShutdownWithContext(ctx)
					},
				})
			}),
	)
}
