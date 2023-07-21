package supplier

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"fmt"
)

const (
	cacheContextName string = "supplier"
	supplierNotFound string = "supplier could not be found"
)

type (
	GetterServiceInterface interface {
		GetAll(request *models.SupplierRequest) *pagination.PagingResponseJSON
		GetById(request *models.SupplierRequest) (*models.SupplierJSON, error)
		GetBySlug(request *models.SupplierRequest) (*models.SupplierJSON, error)
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

func (service *GetterService) GetAll(request *models.SupplierRequest) *pagination.PagingResponseJSON {
	paging := request.Paging

	var page *pagination.PagingResponseJSON

	cacheKey := service.Cache.CacheKeyBuilder(cacheContextName, "GetterService", "GetAll", paging.PagingToString())
	cachedPage, err := service.Cache.GetCache(cacheKey, &pagination.PagingResponseJSON{})
	if err != nil {
		var suppliers []interface{}
		models, maxRow, _ := service.Repository.GetAll(paging)
		if models != nil {
			for _, supplier := range *models {
				suppliers = append(suppliers, supplier.ToJSON())
			}
		}
		response := pagination.NewPagingResponse(request.Paging, suppliers, maxRow)
		page = response.GetPagingResponse()
		if page.MaxRow > 0 {
			service.Cache.SetCache(cacheKey, page)
		}
	} else {
		page = cachedPage.(*pagination.PagingResponseJSON)
	}

	return page
}

func (service *GetterService) GetById(request *models.SupplierRequest) (*models.SupplierJSON, error) {
	requestId := request.Supplier.ID

	var supplier models.SupplierJSON

	if requestId > uint(0) {
		cacheKey := service.Cache.CacheKeyBuilder(cacheContextName, "GetterService", "GetById", fmt.Sprintf("%d", requestId))
		cachedSupplier, err := service.Cache.GetCache(cacheKey, &models.SupplierJSON{})
		if err != nil {
			model, err := service.Repository.GetById(requestId)
			if err == nil {
				if model.ID > 0 {
					supplier = *model.ToJSON()
					service.Cache.SetCache(cacheKey, &supplier)
					return &supplier, nil
				}
			}
		} else {
			return cachedSupplier.(*models.SupplierJSON), nil
		}
	}

	return nil, custom_error.NotFound(supplierNotFound)
}

func (service *GetterService) GetBySlug(request *models.SupplierRequest) (*models.SupplierJSON, error) {
	requestSlug := request.Supplier.SlugName

	var supplier *models.SupplierJSON
	if requestSlug != "" {
		cacheKey := service.Cache.CacheKeyBuilder(cacheContextName, "GetBySlug", requestSlug)
		cachedSupplier, err := service.Cache.GetCache(cacheKey, &models.SupplierJSON{})
		if err != nil {
			model, err := service.Repository.GetBySlug(requestSlug)
			if err == nil {
				if model.ID > 0 {
					supplier = model.ToJSON()
					service.Cache.SetCache(cacheKey, supplier)
					return supplier, nil
				}
			}
		} else {
			return cachedSupplier.(*models.SupplierJSON), nil
		}
	}

	return nil, custom_error.NotFound(supplierNotFound)
}
