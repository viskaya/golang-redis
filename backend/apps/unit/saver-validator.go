package unit

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"

	"github.com/gosimple/slug"
)

type (
	SaverValidatorInterface interface {
		Validate(request *models.UnitRequest) *custom_error.CustomParamRequestError
	}

	SaverValidator struct {
		GetterService GetterServiceInterface
	}
)

func NewSaverValidator(db *db.DBFactory) *SaverValidator {
	getterService := NewGetterService(db)

	return &SaverValidator{getterService}
}

func (s *SaverValidator) Validate(request *models.UnitRequest) *custom_error.CustomParamRequestError {
	err := custom_error.RequestValidator(request)

	if _, ok := err.Messages["name"]; !ok {
		requestBySlug := &models.UnitRequest{
			Auth: request.Auth,
			Unit: models.UnitJSON{SlugName: slug.Make(request.Unit.Name)},
		}
		if _, errSlug := s.GetterService.GetBySlug(requestBySlug); errSlug == nil {
			err.Messages["name"] = "Unit's name is exist. Please choose another name"
		}
	}

	if len(err.Messages) > 0 {
		return err
	}

	return nil
}
