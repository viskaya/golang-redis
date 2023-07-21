package user

import (
	"aegis_test/libs/db"
	"aegis_test/models"
)

type (
	SaverServiceInterface interface {
		Save(request UserSaveAuthRequest) (*models.UserSaveJSON, error)
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

func (service *SaverService) Save(request *models.UserSaveRequest) (*models.UserJSON, error) {
	userToSave := request.User.ToUser()

	if err := service.Validator.Validate(request); err != nil {
		return nil, err
	}

	service.Cache.RemoveCacheForContext(cacheContextName)

	user, err := service.Repository.Save(userToSave)
	if err == nil {
		return user.ToJSON(), nil
	}

	return nil, err
}
