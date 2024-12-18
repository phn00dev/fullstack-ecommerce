package routes

import (
	adminConstructor "eCommerce/internal/domain/admin/constructor"
	sectionConstructor "eCommerce/internal/domain/section/constructor"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	adminApiV1 := app.Group("/api/admin")
	adminsRoute := adminApiV1.Group("/admins")
	adminsRoute.Get("/", adminConstructor.AdminHandler.GetAll)
	adminsRoute.Get("/:adminID", adminConstructor.AdminHandler.GetOne)
	adminsRoute.Post("/", adminConstructor.AdminHandler.Create)
	adminsRoute.Put("/:adminID", adminConstructor.AdminHandler.Update)
	adminsRoute.Delete("/:adminID", adminConstructor.AdminHandler.Delete)

	// section routes
	sectionRoute := adminApiV1.Group("/sections")
	sectionRoute.Get("/", sectionConstructor.SectionHandler.GetAll)
	sectionRoute.Get("/:sectionID", sectionConstructor.SectionHandler.GetOne)
	sectionRoute.Post("/", sectionConstructor.SectionHandler.Create)
	sectionRoute.Put("/:sectionID", sectionConstructor.SectionHandler.Update)
	sectionRoute.Delete("/:sectionID", sectionConstructor.SectionHandler.Delete)

}
