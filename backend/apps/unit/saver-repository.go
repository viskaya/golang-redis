package unit

import (
	"aegis_test/models"

	"gorm.io/gorm"
)

type (
	SaverRepositoryInterface interface {
		Save(unit *models.Unit) (*models.Unit, error)
	}

	SaverRepository struct {
		DB *gorm.DB
	}
)

func NewSaverRepository(db *gorm.DB) *SaverRepository {
	return &SaverRepository{DB: db}
}

func (s *SaverRepository) Save(unit *models.Unit) (*models.Unit, error) {
	err := s.DB.Save(&unit).Error
	if err == nil {
		return unit, nil
	}

	return nil, err
}
