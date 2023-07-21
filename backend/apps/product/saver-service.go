package product

import (
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
)

type (
	SaverServiceInterface interface {
		GetAll(request *models.ProductRequest) *pagination.PagingResponse
		GetById(request *models.ProductRequest) (*models.ProductJSON, error)
		GetBySlug(request *models.ProductRequest) (*models.ProductJSON, error)
	}

	SaverService struct {
		Repository SaverRepositoryInterface
		Validator  SaverValidatorInterface
		Cache      db.CacheDB
	}
)

func NewSaverService(db *db.DBFactory) *SaverService {
	return &SaverService{
		Repository: NewSaverRepository(db.DB),
		Validator:  NewSaverValidator(db),
		Cache:      db.Cache,
	}
}

func (service *SaverService) Save(request *models.ProductRequest) (*models.ProductJSON, error) {
	if err := service.Validator.Validate(request); err != nil {
		return nil, err
	}

	productToSave := request.Product.ToProduct()

	service.Cache.RemoveCacheForContext(CacheContextName)

	product, err := service.Repository.Save(productToSave)
	if err == nil {
		return product.ToJSON(), nil
	}

	return nil, err
}
