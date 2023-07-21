package libs

import (
	"aegis_test/libs/custom_error"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

const (
	HeaderSessionID string = "X-SESSION-ID"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		noAuthPath := []string{"/v1", "/v1/", "/v1/login"}
		isAuthorized := slices.Index(noAuthPath, c.FullPath()) > -1

		if !isAuthorized {
			status := http.StatusUnauthorized
			msg := custom_error.Unauthorized("unauthorized access")
			auths := strings.Split(c.GetHeader("Authorization"), " ")
			if len(auths) == 2 && auths[0] == "Bearer" {
				claim, ok := VerifyJWTToken(auths[1])
				if !ok {
					msg = custom_error.Unauthorized("invalid token")
				} else if claim.ExpiresAt.Before(time.Now()) {
					msg = custom_error.Unauthorized("token expired")
				} else {
					c.Set(HeaderSessionID, claim.SessionID)
					isAuthorized = true
				}
			}

			if !isAuthorized {
				c.JSON(status, gin.H{"error": msg})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
