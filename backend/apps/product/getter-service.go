package product

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"fmt"
)

const (
	CacheContextName string = "product"
	productNotFound  string = "product could not be found"
)

type (
	GetterServiceInterface interface {
		GetAll(request *models.ProductRequest) *pagination.PagingResponseJSON
		GetById(request *models.ProductRequest) (*models.ProductJSON, error)
		GetBySlug(request *models.ProductRequest) (*models.ProductJSON, error)
		GetByRegNumber(request *models.ProductRequest) (*models.ProductJSON, error)
	}

	GetterService struct {
		Repository GetterRepositoryInterface
		Cache      db.CacheDB
	}
)

func NewGetterService(db *db.DBFactory) *GetterService {
	return &GetterService{
		Repository: NewGetterRepository(db.DB),
		Cache:      db.Cache,
	}
}

func (service *GetterService) GetAll(request *models.ProductRequest) *pagination.PagingResponseJSON {
	paging := request.Paging

	var page *pagination.PagingResponseJSON

	cacheKey := service.Cache.CacheKeyBuilder(CacheContextName, "GetterService", "GetAll", paging.PagingToString())
	cachedPage, err := service.Cache.GetCache(cacheKey, &pagination.PagingResponseJSON{})
	if err != nil {
		var products []interface{}
		models, maxRow, _ := service.Repository.GetAll(paging)
		if models != nil {
			for _, product := range *models {
				products = append(products, product.ToJSON())
			}
		}
		response := pagination.NewPagingResponse(request.Paging, products, maxRow)
		page = response.GetPagingResponse()
		if page.MaxRow > 0 {
			service.Cache.SetCache(cacheKey, page)
		}
	} else {
		page = cachedPage.(*pagination.PagingResponseJSON)
	}

	return page
}

func (service *GetterService) GetById(request *models.ProductRequest) (*models.ProductJSON, error) {
	requestId := request.Product.ID

	var product *models.ProductJSON

	if requestId > uint(0) {
		cacheKey := service.Cache.CacheKeyBuilder(CacheContextName, "GetterService", "GetById", fmt.Sprintf("%d", requestId))
		cachedProduct, err := service.Cache.GetCache(cacheKey, &models.ProductJSON{})
		if err != nil {
			model, err := service.Repository.GetById(requestId)
			if err == nil {
				if model.ID > 0 {
					product = model.ToJSON()
					service.Cache.SetCache(cacheKey, &product)
					return product, nil
				}
			}
		} else {
			return cachedProduct.(*models.ProductJSON), nil
		}
	}

	return nil, custom_error.NotFound(productNotFound)
}

func (service *GetterService) GetBySlug(request *models.ProductRequest) (*models.ProductJSON, error) {
	requestSlug := request.Product.SlugName

	var product *models.ProductJSON
	if requestSlug != "" {
		cacheKey := service.Cache.CacheKeyBuilder(CacheContextName, "GetBySlug", requestSlug)
		cachedProduct, err := service.Cache.GetCache(cacheKey, &models.ProductJSON{})
		if err != nil {
			model, err := service.Repository.GetBySlug(requestSlug)
			if err == nil {
				if model.ID > 0 {
					product = model.ToJSON()
					service.Cache.SetCache(cacheKey, product)
					return product, nil
				}
			}
		} else {
			return cachedProduct.(*models.ProductJSON), nil
		}
	}

	return nil, custom_error.NotFound(productNotFound)
}

func (service *GetterService) GetByRegNumber(request *models.ProductRequest) (*models.ProductJSON, error) {
	requestRegNumber := request.Product.RegNumber

	var product *models.ProductJSON
	if requestRegNumber != "" {
		cacheKey := service.Cache.CacheKeyBuilder(CacheContextName, "GetByRegNumber", requestRegNumber)
		cachedProduct, err := service.Cache.GetCache(cacheKey, &models.ProductJSON{})
		if err != nil {
			model, err := service.Repository.GetByRegNumber(requestRegNumber)
			if err == nil {
				if model.ID > 0 {
					product = model.ToJSON()
					service.Cache.SetCache(cacheKey, product)
					return product, nil
				}
			}
		} else {
			return cachedProduct.(*models.ProductJSON), nil
		}
	}

	return nil, custom_error.NotFound(productNotFound)
}
