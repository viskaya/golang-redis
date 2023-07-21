package cashier

import (
	"aegis_test/apps/product"
	"aegis_test/libs"
	"aegis_test/libs/db"
	"aegis_test/models"
)

const (
	cacheContextName string = "cashier"
)

type (
	SaverServiceInterface interface {
		Save(request *models.CashierRequest) (*models.CashierJSON, error)
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

func (service *SaverService) Save(request *models.CashierRequest) (*models.CashierJSON, error) {
	if err := service.Validator.Validate(request); err != nil {
		return nil, err
	}

	service.Cache.RemoveCacheForContext(product.CacheContextName)
	service.Cache.RemoveCacheForContext(cacheContextName)

	cashierToSave := request.Cashier.ToCashier()
	cashierToSave.TrxNumber = libs.GenerateTrxNumber()

	cashier, err := service.Repository.Save(cashierToSave)
	if err == nil {
		return cashier.ToJSON(), nil
	}

	return nil, err
}
