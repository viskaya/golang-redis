package user

import (
	"aegis_test/libs/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetterControllerInterface interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
}

type GetterController struct {
	Auth    UserAuthRequest
	Service GetterService
}

func NewGetterController(db *db.DBFactory) *GetterController {
	auth := NewUserAuthRequest(db)
	service := NewGetterService(db)

	return &GetterController{
		Auth:    *auth,
		Service: *service,
	}
}

func (ctrl *GetterController) GetAll(c *gin.Context) {
	request, err := ctrl.Auth.UserAuthorizeRequest(c)
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
	request, err := ctrl.Auth.UserAuthorizeRequest(c)
	if err == nil {
		user, err := ctrl.Service.GetById(request)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": user,
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
