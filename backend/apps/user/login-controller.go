package user

import (
	"aegis_test/libs/db"
	"aegis_test/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginControllerInterface interface {
	Login(c *gin.Context)
}

type LoginController struct {
	Service LoginServiceInterface
}

func NewLoginController(db *db.DBFactory) *LoginController {
	service := NewLoginService(db)
	return &LoginController{
		Service: service,
	}
}

func (ctrl *LoginController) Login(c *gin.Context) {
	var request models.Login

	c.ShouldBindJSON(&request)

	user, err := ctrl.Service.LoginByUserName(request.UserName, request.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
