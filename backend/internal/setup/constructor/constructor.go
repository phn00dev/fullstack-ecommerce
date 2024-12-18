package constructor

import (
	"eCommerce/internal/app"
	adminConstructor "eCommerce/internal/domain/admin/constructor"
	sectionConstructor "eCommerce/internal/domain/section/constructor"
)

func Build(dependencies app.Dependencies) {
	adminConstructor.InitAdminRequirementCreator(dependencies.DB)
	sectionConstructor.InitSectionRequirementCreator(dependencies.DB, *dependencies.Config)
}
