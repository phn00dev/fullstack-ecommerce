package routes

import (
	adminConstructor "eCommerce/internal/domain/admin/constructor"
	bannerConstructor "eCommerce/internal/domain/banner/constructor"
	brandConstructor "eCommerce/internal/domain/brand/constructor"
	categoryConsturtor "eCommerce/internal/domain/category/consturtor"
	productConstructor "eCommerce/internal/domain/product/constructor"
	sectionConstructor "eCommerce/internal/domain/section/constructor"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	adminApiV1 := app.Group("/api/admin")

	// admin auth routes

	// admin crud routes
	adminsRoute := adminApiV1.Group("/admins")
	adminsRoute.Get("/", adminConstructor.AdminHandler.GetAll)
	adminsRoute.Get("/:adminID", adminConstructor.AdminHandler.GetOne)
	adminsRoute.Post("/", adminConstructor.AdminHandler.Create)
	adminsRoute.Put("/:adminID", adminConstructor.AdminHandler.Update)
	adminsRoute.Delete("/:adminID", adminConstructor.AdminHandler.Delete)

	// banner routes
	bannerRoute := adminApiV1.Group("/banners")
	bannerRoute.Get("/", bannerConstructor.BannerHandler.GetAll)
	bannerRoute.Get("/:bannerID", bannerConstructor.BannerHandler.GetOne)
	bannerRoute.Post("/", bannerConstructor.BannerHandler.Create)
	bannerRoute.Put("/:bannerID", bannerConstructor.BannerHandler.Update)
	bannerRoute.Delete("/:bannerID", bannerConstructor.BannerHandler.Delete)

	// brand routes
	brandRoute := adminApiV1.Group("/brands")
	brandRoute.Get("/", brandConstructor.BrandHandler.GetAll)
	brandRoute.Get("/:brandID", brandConstructor.BrandHandler.GetOne)
	brandRoute.Post("/", brandConstructor.BrandHandler.Create)
	brandRoute.Put("/:brandID", brandConstructor.BrandHandler.Update)
	brandRoute.Delete("/:brandID", brandConstructor.BrandHandler.Delete)

	// section routes
	sectionRoute := adminApiV1.Group("/sections")
	sectionRoute.Get("/", sectionConstructor.SectionHandler.GetAll)
	sectionRoute.Get("/:sectionID", sectionConstructor.SectionHandler.GetOne)
	sectionRoute.Post("/", sectionConstructor.SectionHandler.Create)
	sectionRoute.Put("/:sectionID", sectionConstructor.SectionHandler.Update)
	sectionRoute.Delete("/:sectionID", sectionConstructor.SectionHandler.Delete)

	// category routes

	categoryRoute := adminApiV1.Group("/categories")
	categoryRoute.Get("/", categoryConsturtor.CategoryHandler.GetAll)
	categoryRoute.Get("/:categoryID", categoryConsturtor.CategoryHandler.GetOne)
	categoryRoute.Post("/", categoryConsturtor.CategoryHandler.Create)
	categoryRoute.Put("/:categoryID", categoryConsturtor.CategoryHandler.Update)
	categoryRoute.Delete("/:categoryID", categoryConsturtor.CategoryHandler.Delete)

	//  product routes
	productRoute := adminApiV1.Group("/products")
	productRoute.Get("/", productConstructor.ProductHandler.GetAll)
	productRoute.Get("/:productID", productConstructor.ProductHandler.GetOne)
	productRoute.Post("/", productConstructor.ProductHandler.Create)
	productRoute.Put("/:productID", productConstructor.ProductHandler.Update)
	productRoute.Delete("/:productID", productConstructor.ProductHandler.Delete)

}
