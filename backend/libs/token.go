package libs

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	jwt.RegisteredClaims
	SessionID string
	Roles     []string
}

func CreateJWTToken(sessionId string, sessionExpire time.Time, profiles []string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(sessionExpire),
		},
		SessionID: sessionId,
		Roles:     profiles,
	})

	token, err := claims.SignedString([]byte(HashKey))

	if err != nil {
		log.Println(err)
		token = ""
	}

	return token
}

func VerifyJWTToken(authToken string) (UserClaim, bool) {
	var userClaim UserClaim

	token, err := jwt.ParseWithClaims(authToken, &userClaim, func(t *jwt.Token) (interface{}, error) {
		return []byte(HashKey), nil
	})
	if err != nil {
		log.Println("Error Parse Auth Token: ", err)
		return UserClaim{}, false
	}
	if !token.Valid {
		log.Println("Error Auth Token: Invalid Token")
		return UserClaim{}, false
	}

	return userClaim, true
}
