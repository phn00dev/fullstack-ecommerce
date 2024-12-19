package routes

import (
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func StaticRoutes(app *fiber.App, config *config.Config) {
	app.Static("/public", config.FolderConfig.PublicPath)
}
