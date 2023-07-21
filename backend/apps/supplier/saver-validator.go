package supplier

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"

	"github.com/gosimple/slug"
)

type (
	SaverValidatorInterface interface {
		Validate(request *models.SupplierRequest) *custom_error.CustomParamRequestError
	}

	SaverValidator struct {
		GetterService GetterServiceInterface
	}
)

func NewSaverValidator(db *db.DBFactory) *SaverValidator {
	getterService := NewGetterService(db)

	return &SaverValidator{getterService}
}

func (s *SaverValidator) Validate(request *models.SupplierRequest) *custom_error.CustomParamRequestError {
	err := custom_error.RequestValidator(request)

	if _, ok := err.Messages["name"]; !ok {
		requestBySlug := &models.SupplierRequest{
			Auth:     request.Auth,
			Supplier: models.SupplierJSON{SlugName: slug.Make(request.Supplier.Name)},
		}
		if _, errSlug := s.GetterService.GetBySlug(requestBySlug); errSlug == nil {
			err.Messages["name"] = "Supplier's name is exist. Please choose another name"
		}
	}

	if len(err.Messages) > 0 {
		return err
	}

	return nil
}
