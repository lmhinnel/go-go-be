package server

import (
	"os"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/lmhuong711/go-go-be/routes"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	jwt_secret := "jwt_secret"
	if os.Getenv("jwt_secret") != "" {
		jwt_secret = os.Getenv("jwt_secret")
	}

	app := fiber.New()

	prometheus := fiberprometheus.New("go-go-be")
	prometheus.RegisterAt(app, "/metrics")

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(prometheus.Middleware)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwt_secret)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
				"count":   0,
			})
		},
	}))

	app.Use(func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return nil
	})

	api := app.Group("/api/v1")
	routes.SetupRoutes(api)

	return app
}
