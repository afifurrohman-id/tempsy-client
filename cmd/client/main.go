package main

import (
	"os"
	"path"
	"time"

	"github.com/afifurrohman-id/tempsy-client/pkg/middleware"
	"github.com/afifurrohman-id/tempsy-client/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(path.Join("configs", ".env")); err != nil {
			log.Error(err)
		}
	}
}

func main() {
	engine := html.New(path.Join("web", "template"), ".go.html")
	app := fiber.New(fiber.Config{
		Views:         engine,
		BodyLimit:     30 << 20, // 30MB
		CaseSensitive: true,
		ErrorHandler:  middleware.CatchServerError,
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}), recover.New(), logger.New(), favicon.New(), middleware.Limiter)

	cacheDuration := 10 * time.Second
	if os.Getenv("APP_ENV") != "production" {
		cacheDuration = -1 * time.Second
	}

	app.Static("/public", path.Join("web", "static"), fiber.Static{
		CacheDuration: cacheDuration,
	})

	app.Get("/", middleware.CheckAuth, router.HandleWelcomeClient)

	routeAuth := app.Group("/auth")
	routeAuth.Get("/", router.OAuth2Callback)
	routeAuth.Get("/login", router.AuthLogin)
	routeAuth.Get("/logout", router.AuthLogout)

	routeDashboardUser := app.Group("/dashboard/:username", middleware.CheckAuth)
	routeDashboardUser.Get("/", router.HandleDashboardClient)
	routeDashboardUser.Get("/profile", router.HandleProfileClient)
	routeDashboardUser.Get("/:name", router.HandleDetailDataClient)

	routeDashboardUser.Use(middleware.SetRealIpClient)
	routeDashboardUser.Post("/", router.HandleUploadDataClient)
	routeDashboardUser.Put("/:name", router.HandleUpdateDataClient)
	routeDashboardUser.Delete("/profile", router.HandleDeleteAllDataClient)
	routeDashboardUser.Delete("/:name", router.HandleDeleteDataClient)

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Panic(err)
	}
}
