package service

import (
	brandRepository "eCommerce/internal/domain/brand/repository"
	categoryRepository "eCommerce/internal/domain/category/repository"
	"eCommerce/internal/domain/product/dto"
	"eCommerce/internal/domain/product/repository"
	sectionRepository "eCommerce/internal/domain/section/repository"
	"eCommerce/internal/models"
	"eCommerce/internal/utils/images"
	"eCommerce/pkg/config"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type productServiceImp struct {
	productRepo  repository.ProductRepository
	sectionRepo  sectionRepository.SectionRepository
	categoryRepo categoryRepository.CategoryRepository
	brandRepo    brandRepository.BrandRepository
}

func NewProductService(productRepo repository.ProductRepository,
	sectionRepo sectionRepository.SectionRepository, categoryRepo categoryRepository.CategoryRepository,
	brandRepo brandRepository.BrandRepository) ProductService {
	return productServiceImp{
		productRepo:  productRepo,
		sectionRepo:  sectionRepo,
		categoryRepo: categoryRepo,
		brandRepo:    brandRepo,
	}
}

func (p productServiceImp) GetAllProducts() ([]dto.ProductResponse, error) {
	products, err := p.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	productResponses := dto.GetAllProductResponses(products)
	return productResponses, nil
}

func (p productServiceImp) GetOneProductById(productId int) (*dto.ProductResponse, error) {
	product, err := p.productRepo.GetOne(productId)
	if err != nil {
		return nil, err
	}
	productResponse := dto.GetOneProductResponse(*product)
	return &productResponse, nil
}

func (p productServiceImp) CreateProduct(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateProductRequest) error {
	// check product name tk
	checkProductNameTk, err := p.productRepo.CheckProductNameTk(createRequest.ProductNameTk)
	if err != nil {
		return fmt.Errorf("error checking product name TK: %w", err)
	}
	if checkProductNameTk {
		return errors.New("product name TK already exists")
	}
	// check product name ru
	checkProductNameRu, err := p.productRepo.CheckProductNameTk(createRequest.ProductNameRu)
	if err != nil {
		return fmt.Errorf("error checking product name Ru: %w", err)
	}
	if checkProductNameRu {
		return errors.New("product name Ru already exists")
	}
	// check product name en
	checkProductNameEn, err := p.productRepo.CheckProductNameTk(createRequest.ProductNameEn)
	if err != nil {
		return fmt.Errorf("error checking product name En: %w", err)
	}
	if checkProductNameEn {
		return errors.New("product name En already exists")
	}

	// check section id
	section, errSection := p.sectionRepo.GetOneById(createRequest.SectionID)
	if errSection != nil {
		return errors.New("section not found or section id error")
	}
	// check category id
	category, errCategory := p.categoryRepo.GetOne(createRequest.CategoryID)
	if errCategory != nil {
		return errors.New("category not found or category id error")
	}
	// check brand id
	brand, errBrand := p.brandRepo.GetOneByID(createRequest.BrandID)
	if errBrand != nil {
		return errors.New("brand not found or brand id error")
	}

	// upload product image one
	productMainImageOnePath, errMainImageOne := images.UploadFile(ctx, "product_main_image_one", config.FolderConfig.PublicPath, "product-images")
	if errMainImageOne != nil {
		return errMainImageOne
	}
	// upload product image two
	productMainImageTwoPath, errMainImageTwo := images.UploadFile(ctx, "product_main_image_two", config.FolderConfig.PublicPath, "product-images")
	if errMainImageTwo != nil {
		return errMainImageTwo
	}

	newProduct := models.Product{
		ProductNameTk:             createRequest.ProductNameTk,
		ProductNameRu:             createRequest.ProductNameRu,
		ProductNameEn:             createRequest.ProductNameEn,
		ProductShortDescriptionTk: createRequest.ProductShortDescriptionTk,
		ProductShortDescriptionRu: createRequest.ProductShortDescriptionRu,
		ProductShortDescriptionEn: createRequest.ProductShortDescriptionEn,
		ProductLongDescriptionTk:  createRequest.ProductLongDescriptionTk,
		ProductLongDescriptionRu:  createRequest.ProductLongDescriptionRu,
		ProductLongDescriptionEn:  createRequest.ProductLongDescriptionEn,
		ProductAllSpecificationTk: createRequest.ProductAllSpecificationTk,
		ProductAllSpecificationRu: createRequest.ProductAllSpecificationRu,
		ProductAllSpecificationEn: createRequest.ProductAllSpecificationEn,
		ProductPrice:              createRequest.ProductPrice,
		ProductTotalCount:         createRequest.ProductTotalCount,
		ProductRemainingNumber:    createRequest.ProductRemainingNumber,
		ProductMainImageOne:       *productMainImageOnePath,
		ProductMainImageTwo:       *productMainImageTwoPath,
		ProductStatus:             createRequest.ProductStatus,
		SectionID:                 section.ID,
		CategoryID:                category.ID,
		BrandID:                   brand.ID,
	}
	// create product
	if errCreateProduct := p.productRepo.Create(newProduct); errCreateProduct != nil {
		// eger-de product doredilende error bolsa onda  onda upload edilen product imageleri delete etmeli
		// delete one product
		if errDeleteOneImage := images.DeleteFile(*productMainImageOnePath); errDeleteOneImage != nil {
			return errDeleteOneImage
		}
		// delete two product
		if errDeleteTwoImage := images.DeleteFile(*productMainImageOnePath); errDeleteTwoImage != nil {
			return errDeleteTwoImage
		}
		return errCreateProduct
	}
	return nil
}

func (p productServiceImp) UpdateProduct(ctx *fiber.Ctx, config config.Config, productID int, updateRequest dto.UpdateProductRequest) error {
	// get product
	product, err := p.productRepo.GetOne(productID)
	if err != nil {
		return err
	}
	// check section id
	section, errSection := p.sectionRepo.GetOneById(updateRequest.SectionID)
	if errSection != nil {
		return errors.New("section not found or section id error")
	}
	// check category id
	category, errCategory := p.categoryRepo.GetOne(updateRequest.CategoryID)
	if errCategory != nil {
		return errors.New("category not found or category id error")
	}
	// check brand id
	brand, errBrand := p.brandRepo.GetOneByID(updateRequest.BrandID)
	if errBrand != nil {
		return errors.New("brand not found or brand id error")
	}

	// check product main image one
	imageOne, _ := ctx.FormFile("product_main_image_one")
	if imageOne != nil {
		// delete product one image
		if errDeleteOneImage := images.DeleteFile(product.ProductMainImageOne); errDeleteOneImage != nil {
			return errDeleteOneImage
		}
		// upload new product main one image
		productMainImageOnePath, errMainImageOne := images.UploadFile(ctx, "product_main_image_one", config.FolderConfig.PublicPath, "product-images")
		if errMainImageOne != nil {
			return errMainImageOne
		}
		product.ProductMainImageOne = *productMainImageOnePath
	}

	imageTwo, _ := ctx.FormFile("product_main_image_one")
	if imageTwo != nil {
		// delete product one image
		if errDeleteTwoImage := images.DeleteFile(product.ProductMainImageTwo); errDeleteTwoImage != nil {
			return errDeleteTwoImage
		}
		// upload new product main one image
		productMainImageTwoPath, errMainImageTwo := images.UploadFile(ctx, "product_main_image_two", config.FolderConfig.PublicPath, "product-images")
		if errMainImageTwo != nil {
			return errMainImageTwo
		}
		product.ProductMainImageTwo = *productMainImageTwoPath
	}

	product.ProductNameTk = updateRequest.ProductNameTk
	product.ProductNameRu = updateRequest.ProductNameTk
	product.ProductNameEn = updateRequest.ProductNameEn
	product.ProductShortDescriptionTk = updateRequest.ProductShortDescriptionTk
	product.ProductShortDescriptionTk = updateRequest.ProductShortDescriptionTk
	product.ProductShortDescriptionTk = updateRequest.ProductShortDescriptionTk
	product.ProductLongDescriptionTk = updateRequest.ProductLongDescriptionTk
	product.ProductLongDescriptionRu = updateRequest.ProductLongDescriptionRu
	product.ProductLongDescriptionEn = updateRequest.ProductLongDescriptionEn
	product.ProductAllSpecificationTk = updateRequest.ProductAllSpecificationTk
	product.ProductAllSpecificationRu = updateRequest.ProductAllSpecificationRu
	product.ProductAllSpecificationEn = updateRequest.ProductAllSpecificationEn
	product.ProductPrice = updateRequest.ProductPrice
	product.ProductTotalCount = updateRequest.ProductTotalCount
	product.ProductRemainingNumber = updateRequest.ProductRemainingNumber
	product.ProductStatus = updateRequest.ProductStatus
	product.SectionID = section.ID
	product.CategoryID = category.ID
	product.BrandID = brand.ID
	return p.productRepo.Update(int(product.ID), *product)
}

func (p productServiceImp) DeleteProduct(productID int) error {
	// get product
	product, err := p.productRepo.GetOne(productID)
	if err != nil {
		return errors.New("something went wrong. Product not found")
	}
	// delete product one image
	if errDeleteOneProductImage := images.DeleteFile(product.ProductMainImageOne); errDeleteOneProductImage != nil {
		return errDeleteOneProductImage
	}
	// delete product two image
	if errDeleteTwoProductImage := images.DeleteFile(product.ProductMainImageTwo); errDeleteTwoProductImage != nil {
		return errDeleteTwoProductImage
	}
	return p.productRepo.Delete(int(product.ID))
}
