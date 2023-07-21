package auth

import (
	"aegis_test/libs"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	cacheContextName string = "auth"
)

type AuthorizationServiceInterface interface {
	AuthorizeRequest(c *gin.Context) (models.AuthorizedRequest, error)
}

type AuthorizationRequest struct {
	Repository AuthRepository
	Cache      db.CacheDB
}

func NewAuthorizationRequest(db *gorm.DB, cache db.CacheDB) *AuthorizationRequest {
	return &AuthorizationRequest{
		Repository: *NewAuthRepository(db),
		Cache:      cache,
	}
}

func (auth *AuthorizationRequest) AuthorizeRequest(c *gin.Context) (*models.AuthorizedRequest, error) {
	if sessionId, ok := c.Get(libs.HeaderSessionID); ok {
		cacheKey := auth.Cache.CacheKeyBuilder(cacheContextName, fmt.Sprintf("%v", sessionId))
		userCache, err := auth.Cache.GetCache(cacheKey, &models.AuthorizedUser{})
		if err != nil {
			user, err := auth.Repository.GetUserBySessionId(sessionId.(string))
			if err == nil && user.ID > 0 {
				authUser := user.ToAuthorizedUser()
				auth.Cache.SetCache(cacheKey, authUser)
				return &models.AuthorizedRequest{
					User:    authUser,
					Request: c,
				}, nil
			}
		} else {
			return &models.AuthorizedRequest{
				User:    userCache.(*models.AuthorizedUser),
				Request: c,
			}, nil
		}
	}

	return nil, custom_error.Unauthorized("invalid token or token expired")
}
