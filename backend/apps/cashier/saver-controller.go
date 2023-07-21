package cashier

import (
	"aegis_test/libs/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaverControllerInterface interface {
	GetAll(*gin.Context)
}

type SaverController struct {
	Auth    CashierAuthRequest
	Service SaverService
}

func NewSaverController(db *db.DBFactory) *SaverController {
	auth := NewCashierAuthRequest(db)
	service := NewSaverService(db)

	return &SaverController{
		Auth:    *auth,
		Service: *service,
	}
}

func (ctrl *SaverController) Save(c *gin.Context) {
	request, err := ctrl.Auth.CashierAuthorizeRequest(c)
	if err == nil {
		cashier, err := ctrl.Service.Save(request)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": cashier,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}
}
