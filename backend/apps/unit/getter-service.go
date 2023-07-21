package unit

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"fmt"
)

type (
	GetterServiceInterface interface {
		GetAll(request *models.UnitRequest) *pagination.PagingResponseJSON
		GetById(request *models.UnitRequest) (*models.UnitJSON, error)
		GetBySlug(request *models.UnitRequest) (*models.UnitJSON, error)
	}

	GetterService struct {
		Repository GetterRepositoryInterface
		Cache      db.CacheDB
	}
)

const (
	cacheContextName    string = "unit"
	unitCouldNotBeFound string = "unit could not be found"
)

func NewGetterService(db *db.DBFactory) *GetterService {
	repo := NewGetterRepository(db.DB)
	return &GetterService{
		Repository: repo,
		Cache:      db.Cache,
	}
}

func (service *GetterService) GetAll(request *models.UnitRequest) *pagination.PagingResponseJSON {
	paging := request.Paging

	var page *pagination.PagingResponseJSON

	cacheKey := service.Cache.CacheKeyBuilder(cacheContextName, "GetterService", "GetAll", paging.PagingToString())
	cachedPage, err := service.Cache.GetCache(cacheKey, &pagination.PagingResponseJSON{})
	if err != nil {
		var units []interface{}
		models, maxRow, _ := service.Repository.GetAll(paging)
		if models != nil {
			for _, unit := range *models {
				units = append(units, unit.ToJSON())
			}
		}
		response := pagination.NewPagingResponse(paging, units, maxRow)
		page = response.GetPagingResponse()
		if page.MaxRow > 0 {
			service.Cache.SetCache(cacheKey, page)
		}
	} else {
		page = cachedPage.(*pagination.PagingResponseJSON)
	}

	return page
}

func (service *GetterService) GetById(request *models.UnitRequest) (*models.UnitJSON, error) {
	unitId := request.Unit.ID

	var unit *models.UnitJSON

	if unitId > 0 {
		cacheKey := service.Cache.CacheKeyBuilder(cacheContextName, "GetById", fmt.Sprintf("%d", unitId))
		requestUnit, err := service.Cache.GetCache(cacheKey, &models.UnitJSON{})
		if err != nil {
			model, err := service.Repository.GetById(unitId)
			if err == nil && model.ID > 0 {
				unit = model.ToJSON()
				service.Cache.SetCache(cacheKey, unit)
				return unit, nil
			} else {
				return nil, custom_error.NotFound(unitCouldNotBeFound)
			}
		} else {
			return requestUnit.(*models.UnitJSON), nil
		}
	}

	return nil, custom_error.NotFound(unitCouldNotBeFound)
}

func (service *GetterService) GetBySlug(request *models.UnitRequest) (*models.UnitJSON, error) {
	unitSlug := request.Unit.SlugName

	var unit *models.UnitJSON

	if unitSlug != "" {
		cacheKey := service.Cache.CacheKeyBuilder(cacheContextName, "GetBySlug", unitSlug)
		cachedUnit, err := service.Cache.GetCache(cacheKey, &models.UnitJSON{})
		if err != nil {
			model, err := service.Repository.GetBySlug(unitSlug)
			if err == nil && model.ID > 0 {
				unit = model.ToJSON()
				service.Cache.SetCache(cacheKey, unit)
				return unit, nil
			}
		} else {
			return cachedUnit.(*models.UnitJSON), nil
		}
	}

	return nil, custom_error.NotFound(unitCouldNotBeFound)
}
