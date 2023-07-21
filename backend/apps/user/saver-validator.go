package user

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"
)

type (
	SaverValidatorInterface interface {
		Validate(request *models.UserSaveRequest) *custom_error.CustomParamRequestError
	}

	SaverValidator struct {
		GetterService GetterServiceInterface
	}
)

func NewSaverValidator(db *db.DBFactory) *SaverValidator {
	getterService := NewGetterService(db)

	return &SaverValidator{getterService}
}

func (s *SaverValidator) Validate(request *models.UserSaveRequest) *custom_error.CustomParamRequestError {
	invalid := custom_error.RequestValidator(request)

	s.validateUserName(request, invalid)

	if len(invalid.Messages) > 0 {
		return invalid
	}

	return nil
}

func (s *SaverValidator) validateUserName(request *models.UserSaveRequest, invalid *custom_error.CustomParamRequestError) {
	if _, ok := invalid.Messages["name"]; !ok {
		requestByUserName := &models.UserRequest{
			Auth: request.Auth,
			User: models.UserJSON{UserName: request.User.UserName},
		}
		if _, errSlug := s.GetterService.GetByUserName(requestByUserName); errSlug == nil {
			invalid.Messages["userName"] = "User Name is exist. Please choose another"
		}
	}
}
