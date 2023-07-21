package supplier

import (
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
)

type (
	SaverServiceInterface interface {
		GetAll(request *models.SupplierRequest) *pagination.PagingResponse
		GetById(request *models.SupplierRequest) (*models.SupplierJSON, error)
		GetBySlug(request *models.SupplierRequest) (*models.SupplierJSON, error)
	}

	SaverService struct {
		Repository SaverRepositoryInterface
		Validator  SaverValidatorInterface
		Cache      db.CacheDB
	}
)

func NewSaverService(db *db.DBFactory) *SaverService {
	repo := NewSaverRepository(db.DB)
	validator := NewSaverValidator(db)
	return &SaverService{
		Repository: repo,
		Validator:  validator,
		Cache:      db.Cache,
	}
}

func (service *SaverService) Save(request *models.SupplierRequest) (*models.SupplierJSON, error) {
	supplierToSave := request.Supplier.ToSupplier()

	if err := service.Validator.Validate(request); err != nil {
		return nil, err
	}

	service.Cache.RemoveCacheForContext(cacheContextName)

	supplier, err := service.Repository.Save(&supplierToSave)
	if err == nil {
		return supplier.ToJSON(), nil
	}

	return nil, err
}
