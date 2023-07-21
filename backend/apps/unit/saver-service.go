package unit

import (
	"aegis_test/libs/db"
	"aegis_test/models"
)

type (
	SaverServiceInterface interface {
		Save(request UnitAuthRequest) (*models.UnitJSON, error)
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

func (service *SaverService) Save(request *models.UnitRequest) (*models.UnitJSON, error) {
	unitToSave := request.Unit.ToUnit()

	if err := service.Validator.Validate(request); err != nil {
		return nil, err
	}

	service.Cache.RemoveCacheForContext(cacheContextName)

	respUnit, err := service.Repository.Save(&unitToSave)
	if err == nil {
		return respUnit.ToJSON(), nil
	}

	return nil, err
}
