package supplier

import (
	"aegis_test/models"

	"gorm.io/gorm"
)

type (
	SaverRepositoryInterface interface {
		Save(unit *models.Supplier) (*models.Supplier, error)
	}

	SaverRepository struct {
		DB *gorm.DB
	}
)

func NewSaverRepository(db *gorm.DB) *SaverRepository {
	return &SaverRepository{DB: db}
}

func (s *SaverRepository) Save(supplier *models.Supplier) (*models.Supplier, error) {
	err := s.DB.Save(&supplier).Error
	if err == nil {
		return supplier, nil
	}

	return nil, err
}
