package models

import (
	"github.com/gin-gonic/gin"
)

type AuthorizedRequest struct {
	User    *AuthorizedUser
	Request *gin.Context
}
