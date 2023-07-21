package cashier

import (
	"aegis_test/apps/auth"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CashierAuthRequestInterface interface {
	CashierAuthorizeRequest(c *gin.Context) (*models.CashierRequest, error)
}

type CashierAuthRequest struct {
	Auth *auth.AuthorizationRequest
}

func NewCashierAuthRequest(db *db.DBFactory) *CashierAuthRequest {
	auth := auth.NewAuthorizationRequest(db.DB, db.Cache)
	return &CashierAuthRequest{auth}
}

func (auth *CashierAuthRequest) CashierAuthorizeRequest(c *gin.Context) (*models.CashierRequest, error) {
	request := models.CashierRequest{}

	authReq, err := auth.Auth.AuthorizeRequest(c)
	if err == nil {
		var cashier *models.CashierJSON

		if id, err := strconv.Atoi(c.Param("id")); err == nil {
			cashier.ID = uint(id)
		} else {
			err = c.ShouldBindJSON(&cashier)
		}

		if err == nil {
			request = models.CashierRequest{
				Auth:    authReq,
				Cashier: cashier,
			}
		} else {
			err = custom_error.BadRequest("Invalid Request")
		}
	}

	return &request, err
}
