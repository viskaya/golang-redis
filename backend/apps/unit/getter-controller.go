package unit

import (
	"aegis_test/libs/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetterControllerInterface interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	GetBySlug(*gin.Context)
}

type GetterController struct {
	Auth    UnitAuthRequest
	Service GetterService
}

func NewGetterController(db *db.DBFactory) *GetterController {
	auth := NewUnitAuthRequest(db)
	service := NewGetterService(db)

	return &GetterController{
		Auth:    *auth,
		Service: *service,
	}
}

func (ctrl *GetterController) GetAll(c *gin.Context) {
	request, err := ctrl.Auth.UnitAuthorizeRequest(c)
	if err == nil {
		page := ctrl.Service.GetAll(request)
		c.JSON(http.StatusOK, gin.H{
			"data": page,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}
}

func (ctrl *GetterController) GetById(c *gin.Context) {
	request, err := ctrl.Auth.UnitAuthorizeRequest(c)
	if err == nil {
		unit, err := ctrl.Service.GetById(request)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": unit,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

}

func (ctrl *GetterController) GetBySlug(c *gin.Context) {
	request, err := ctrl.Auth.UnitAuthorizeRequest(c)
	if err == nil {
		unit, err := ctrl.Service.GetBySlug(request)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": unit,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}
}
