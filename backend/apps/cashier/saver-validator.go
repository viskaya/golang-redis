package cashier

import (
	"aegis_test/apps/product"
	"aegis_test/apps/unit"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"
	"fmt"
)

type (
	SaverValidatorInterface interface {
		Validate(request *models.CashierRequest) *custom_error.CustomParamRequestError
	}

	SaverValidator struct {
		ProductService product.GetterServiceInterface
		UnitService    unit.GetterServiceInterface
	}
)

func NewSaverValidator(db *db.DBFactory) *SaverValidator {

	return &SaverValidator{
		ProductService: product.NewGetterService(db),
		UnitService:    unit.NewGetterService(db),
	}
}

func (s *SaverValidator) Validate(request *models.CashierRequest) *custom_error.CustomParamRequestError {
	invalid := custom_error.RequestValidator(request)

	s.validateProduct(request, invalid)

	if len(invalid.Messages) > 0 {
		return invalid
	}

	return nil
}

func (s *SaverValidator) validateProduct(request *models.CashierRequest, invalid *custom_error.CustomParamRequestError) {
	if _, ok := invalid.Messages["productID"]; !ok {
		requestById := &models.ProductRequest{
			Auth:    request.Auth,
			Product: models.ProductJSON{ID: request.Cashier.ProductID},
		}
		eq, err := s.ProductService.GetById(requestById)
		if err != nil {
			invalid.Messages["productID"] = "invalid product"
		} else if eq.Quantity < request.Cashier.Quantity {
			invalid.Messages["productID"] = "insufficient quantity"
		} else {
			payment := request.Cashier.Quantity * eq.SellingPrice
			if request.Cashier.TotalPayment < payment {
				invalid.Messages["productID"] = "payment must be greater or equal " + fmt.Sprintf("%2.f", payment)
			}
		}
	}
}
