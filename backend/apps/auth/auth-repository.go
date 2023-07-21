package auth

import (
	"aegis_test/models"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	GetUserBySessionId(sessionId string) (*models.User, error)
}

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (repo *AuthRepository) GetUserBySessionId(sessionId string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("session_id = ?", sessionId).First(&user).Error
	if err == nil {
		return &user, nil
	}

	return nil, err
}
