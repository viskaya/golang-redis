package product

import (
	"aegis_test/apps/auth"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductAuthRequestInterface interface {
	ProductAuthorizeRequest(c *gin.Context) (*models.ProductRequest, error)
}

type ProductAuthRequest struct {
	Auth *auth.AuthorizationRequest
}

func NewProductAuthRequest(db *db.DBFactory) *ProductAuthRequest {
	auth := auth.NewAuthorizationRequest(db.DB, db.Cache)
	return &ProductAuthRequest{auth}
}

func (auth *ProductAuthRequest) ProductAuthorizeRequest(c *gin.Context) (*models.ProductRequest, error) {
	request := models.ProductRequest{}

	authReq, err := auth.Auth.AuthorizeRequest(c)
	if err == nil {
		var product models.ProductJSON

		if id, err := strconv.Atoi(c.Param("id")); err == nil {
			product.ID = uint(id)
		} else if slug := c.Param("slug"); slug != "" {
			err = nil
			product.SlugName = string(slug)
		} else {
			err = c.ShouldBindJSON(&product)
		}

		if err == nil {
			sortFields := make(map[string]string)

			request = models.ProductRequest{
				Auth:    authReq,
				Paging:  pagination.NewPagingRequest(c, sortFields),
				Product: product,
			}
		} else {
			err = custom_error.BadRequest("Invalid Request")
		}
	}

	return &request, err
}
