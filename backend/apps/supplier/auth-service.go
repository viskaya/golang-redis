package supplier

import (
	"aegis_test/apps/auth"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierAuthRequestInterface interface {
	SupplierAuthorizeRequest(c *gin.Context) (*models.SupplierRequest, error)
}

type SupplierAuthRequest struct {
	Auth *auth.AuthorizationRequest
}

func NewSupplierAuthRequest(auth *auth.AuthorizationRequest) *SupplierAuthRequest {
	return &SupplierAuthRequest{auth}
}

func (auth *SupplierAuthRequest) SupplierAuthorizeRequest(c *gin.Context) (*models.SupplierRequest, error) {
	request := models.SupplierRequest{}

	authReq, err := auth.Auth.AuthorizeRequest(c)
	if err == nil {
		var supplier models.SupplierJSON

		if id, err := strconv.Atoi(c.Param("id")); err == nil {
			supplier.ID = uint(id)
		} else if slug := c.Param("slug"); slug != "" {
			err = nil
			supplier.SlugName = string(slug)
		} else {
			err = c.ShouldBindJSON(&supplier)
			log.Printf("bind json %v, error = %v", supplier, err)
		}

		if err == nil {
			sortFields := make(map[string]string)
			sortFields["name"] = "name"

			request = models.SupplierRequest{
				Auth:     authReq,
				Paging:   pagination.NewPagingRequest(c, sortFields),
				Supplier: supplier,
			}
		} else {
			err = custom_error.BadRequest("Invalid Request")
		}
	}

	return &request, err
}
