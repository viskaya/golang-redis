package user

import (
	"aegis_test/libs/db"
	"aegis_test/models"
	"time"

	"gorm.io/gorm"
)

type LoginRepositoryInterface interface {
	GetByUserName(userName string) (models.User, error)
	UpdateAfterLogin(user *models.User) error
}

type LoginRepository struct {
	DB *gorm.DB
}

func NewLoginRepository(db *db.DBFactory) *LoginRepository {
	return &LoginRepository{db.DB}
}

func (repo *LoginRepository) GetByUserName(userName string) (models.User, error) {
	var userModel models.User
	err := repo.DB.Preload("UserProfiles.Profile").
		Where("email = ?", userName).
		Find(&userModel).
		Error

	return userModel, err
}

func (repo *LoginRepository) UpdateAfterLogin(user *models.User) error {
	updateUser := models.User{
		LastLoginAt:     time.Now(),
		SessionID:       user.SessionID,
		SessionExpireAt: user.SessionExpireAt,
	}
	err := repo.DB.Model(&user).Updates(updateUser).Error

	return err
}
