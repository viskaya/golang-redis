package user

import (
	"aegis_test/apps/auth"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"

	"github.com/gin-gonic/gin"
)

type UserSaveAuthRequestInterface interface {
	UserAuthorizeRequest(c *gin.Context) (*models.UserRequest, error)
}

type UserSaveAuthRequest struct {
	Auth *auth.AuthorizationRequest
}

func NewUserSaveAuthRequest(db *db.DBFactory) *UserSaveAuthRequest {
	return &UserSaveAuthRequest{auth.NewAuthorizationRequest(db.DB, db.Cache)}
}

func (auth *UserSaveAuthRequest) UserAuthorizeRequest(c *gin.Context) (*models.UserSaveRequest, error) {
	var request models.UserSaveRequest

	authReq, err := auth.Auth.AuthorizeRequest(c)
	if err == nil {
		var user models.UserSaveJSON

		err = c.ShouldBind(&user)

		if err == nil {
			request = models.UserSaveRequest{
				Auth: authReq,
				User: user,
			}
		} else {
			err = custom_error.BadRequest("Invalid Request")
		}
	}

	return &request, err
}
