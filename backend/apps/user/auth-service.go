package user

import (
	"aegis_test/apps/auth"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAuthRequestInterface interface {
	UserAuthorizeRequest(c *gin.Context) (*models.UserRequest, error)
}

type UserAuthRequest struct {
	Auth *auth.AuthorizationRequest
}

func NewUserAuthRequest(db *db.DBFactory) *UserAuthRequest {
	return &UserAuthRequest{auth.NewAuthorizationRequest(db.DB, db.Cache)}
}

func (auth *UserAuthRequest) UserAuthorizeRequest(c *gin.Context) (*models.UserRequest, error) {
	var request models.UserRequest

	authReq, err := auth.Auth.AuthorizeRequest(c)
	if err == nil {
		var user models.UserJSON

		if id, err := strconv.Atoi(c.Param("id")); err == nil {
			user.ID = uint(id)
		} else {
			err = c.ShouldBind(&user)
		}

		if err == nil {
			sortFields := make(map[string]string)
			sortFields["userName"] = "email"

			request = models.UserRequest{
				Auth:   authReq,
				Paging: pagination.NewPagingRequest(c, sortFields),
				User:   user,
			}
		} else {
			err = custom_error.BadRequest("Invalid Request")
		}
	}

	return &request, err
}
