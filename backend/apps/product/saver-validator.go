package product

import (
	"aegis_test/apps/supplier"
	"aegis_test/apps/unit"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"

	"github.com/gosimple/slug"
)

type (
	SaverValidatorInterface interface {
		Validate(request *models.ProductRequest) *custom_error.CustomParamRequestError
	}

	SaverValidator struct {
		GetterService   GetterServiceInterface
		SupplierService supplier.GetterServiceInterface
		UnitService     unit.GetterServiceInterface
	}
)

func NewSaverValidator(db *db.DBFactory) *SaverValidator {

	return &SaverValidator{
		GetterService:   NewGetterService(db),
		SupplierService: supplier.NewGetterService(db),
		UnitService:     unit.NewGetterService(db),
	}
}

func (s *SaverValidator) Validate(request *models.ProductRequest) *custom_error.CustomParamRequestError {
	invalid := custom_error.RequestValidator(request)

	s.validateRegNumber(request, invalid)
	s.validateName(request, invalid)
	s.validateSupplier(request, invalid)
	s.validateUnit(request, invalid)

	if len(invalid.Messages) > 0 {
		return invalid
	}

	return nil
}

func (s *SaverValidator) validateRegNumber(request *models.ProductRequest, invalid *custom_error.CustomParamRequestError) {
	if _, ok := invalid.Messages["regNumber"]; !ok {
		requestByRegNumber := &models.ProductRequest{
			Auth:    request.Auth,
			Product: models.ProductJSON{RegNumber: request.Product.RegNumber},
		}
		if eq, err := s.GetterService.GetByRegNumber(requestByRegNumber); err == nil && (eq != nil && eq.ID != request.Product.ID) {
			invalid.Messages["regNumber"] = "Registration Number is exist. Please choose another"
		}
	}
}

func (s *SaverValidator) validateName(request *models.ProductRequest, invalid *custom_error.CustomParamRequestError) {
	if _, ok := invalid.Messages["name"]; !ok {
		requestBySlug := &models.ProductRequest{
			Auth:    request.Auth,
			Product: models.ProductJSON{SlugName: slug.Make(request.Product.Name)},
		}
		if eq, err := s.GetterService.GetBySlug(requestBySlug); err == nil && (eq != nil && eq.ID != request.Product.ID) {
			invalid.Messages["name"] = "Product's name is exist. Please choose another name"
		}
	}
}

func (s *SaverValidator) validateSupplier(request *models.ProductRequest, invalid *custom_error.CustomParamRequestError) {
	if request.Product.SupplierID > 0 {
		supplierRequest := &models.SupplierRequest{
			Auth:     request.Auth,
			Supplier: models.SupplierJSON{ID: request.Product.SupplierID},
		}
		if _, err := s.SupplierService.GetById(supplierRequest); err != nil {
			invalid.Messages["supplierID"] = "invalid supplier"
		}
	}
}

func (s *SaverValidator) validateUnit(request *models.ProductRequest, invalid *custom_error.CustomParamRequestError) {
	if request.Product.UnitID > 0 {
		unitRequest := &models.UnitRequest{
			Auth: request.Auth,
			Unit: models.UnitJSON{ID: request.Product.UnitID},
		}
		if _, err := s.UnitService.GetById(unitRequest); err != nil {
			invalid.Messages["unitID"] = "invalid unit"
		}
	}
}
