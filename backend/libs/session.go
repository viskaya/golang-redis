package libs

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserSessionGeneratorInterface interface {
	SessionGenerate(id uint) (*string, *time.Time)
}

type UserSessionGenerator struct {
}

func NewUserSessionGenerator() *UserSessionGenerator {
	return &UserSessionGenerator{}
}

func (gen *UserSessionGenerator) SessionGenerate(id uint) (*string, *time.Time) {
	expire := time.Now().Add(time.Second * time.Duration(TokenTtlHour))
	sessionId, _ := HashPassword(fmt.Sprintf("%d%s", id, jwt.NewNumericDate(expire)))

	return &sessionId, &expire
}
