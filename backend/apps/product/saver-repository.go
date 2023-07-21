package product

import (
	"aegis_test/models"

	"gorm.io/gorm"
)

type (
	SaverRepositoryInterface interface {
		Save(product *models.Products) (*models.Products, error)
	}

	SaverRepository struct {
		DB *gorm.DB
	}
)

func NewSaverRepository(db *gorm.DB) *SaverRepository {
	return &SaverRepository{DB: db}
}

func (s *SaverRepository) Save(product *models.Products) (*models.Products, error) {
	err := s.DB.Save(&product).Error
	if err == nil {
		return product, nil
	}

	return nil, err
}
