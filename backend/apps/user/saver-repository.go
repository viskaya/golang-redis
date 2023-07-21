package user

import (
	"aegis_test/models"

	"gorm.io/gorm"
)

type (
	SaverRepositoryInterface interface {
		Save(user *models.User) (*models.User, error)
	}

	SaverRepository struct {
		DB *gorm.DB
	}
)

func NewSaverRepository(db *gorm.DB) *SaverRepository {
	return &SaverRepository{DB: db}
}

func (s *SaverRepository) Save(user *models.User) (*models.User, error) {
	user.Status = models.UserActive

	err := s.DB.Save(&user).Error
	if err == nil {
		return user, nil
	}

	return nil, err
}
