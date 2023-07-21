package unit

import (
	"aegis_test/apps/auth"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UnitAuthRequestInterface interface {
	UnitAuthorizeRequest(c *gin.Context) (*models.UnitRequest, error)
}

type UnitAuthRequest struct {
	Auth *auth.AuthorizationRequest
}

func NewUnitAuthRequest(db *db.DBFactory) *UnitAuthRequest {
	return &UnitAuthRequest{auth.NewAuthorizationRequest(db.DB, db.Cache)}
}

func (auth *UnitAuthRequest) UnitAuthorizeRequest(c *gin.Context) (*models.UnitRequest, error) {
	var request models.UnitRequest

	authRequest, err := auth.Auth.AuthorizeRequest(c)
	if err == nil {
		var unit models.UnitJSON

		if id, err := strconv.Atoi(c.Param("id")); err == nil {
			unit.ID = uint(id)
		} else if slug := c.Param("slugName"); slug != "" {
			unit.SlugName = slug
			err = nil
		} else {
			err = c.ShouldBind(&unit)
		}

		if err == nil {
			sortFields := make(map[string]string)
			sortFields["name"] = "name"

			request = models.UnitRequest{
				Auth:   authRequest,
				Paging: pagination.NewPagingRequest(c, sortFields),
				Unit:   unit,
			}
		} else {
			err = custom_error.BadRequest("Invalid Request")
		}
	}

	return &request, err
}
