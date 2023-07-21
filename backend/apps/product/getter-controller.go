package product

import (
	"aegis_test/libs/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetterControllerInterface interface {
	GetAll(*gin.Context)
	GetById(c *gin.Context)
	GetBySlug(c *gin.Context)
	GetByRegNumber(c *gin.Context)
}

type GetterController struct {
	Auth    *ProductAuthRequest
	Service GetterService
}

func NewGetterController(db *db.DBFactory) *GetterController {
	auth := NewProductAuthRequest(db)
	service := NewGetterService(db)

	return &GetterController{
		Auth:    auth,
		Service: *service,
	}
}

func (ctrl *GetterController) GetAll(c *gin.Context) {
	request, err := ctrl.Auth.ProductAuthorizeRequest(c)
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
	request, err := ctrl.Auth.ProductAuthorizeRequest(c)
	if err == nil {
		product, err := ctrl.Service.GetById(request)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": product,
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
	request, err := ctrl.Auth.ProductAuthorizeRequest(c)
	if err == nil {
		product, err := ctrl.Service.GetBySlug(request)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": product,
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

func (ctrl *GetterController) GetByRegNumber(c *gin.Context) {
	request, err := ctrl.Auth.ProductAuthorizeRequest(c)
	if err == nil {
		product, err := ctrl.Service.GetByRegNumber(request)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": product,
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
