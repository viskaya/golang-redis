package cashier

import (
	"aegis_test/models"

	"gorm.io/gorm"
)

type (
	SaverRepositoryInterface interface {
		Save(unit *models.Cashiers) (*models.Cashiers, error)
	}

	SaverRepository struct {
		DB *gorm.DB
	}
)

func NewSaverRepository(db *gorm.DB) *SaverRepository {
	return &SaverRepository{DB: db}
}

func (s *SaverRepository) Save(cashier *models.Cashiers) (*models.Cashiers, error) {
	tx := s.DB.Begin()

	tx.Exec("update products set quantity = quantity - ? where id = ?", cashier.Quantity, cashier.ProductID)

	tx.Save(&cashier)

	err := tx.Commit().Error

	if err == nil {
		return cashier, nil
	}

	return nil, err
}
