package product

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/pagination"
	"aegis_test/models"

	"gorm.io/gorm"
)

type GetterRepositoryInterface interface {
	GetAll(paging *pagination.PagingRequest) (*[]models.Products, int64, error)
	GetById(id uint) (*models.Products, error)
	GetBySlug(slug string) (*models.Products, error)
	GetByRegNumber(regNumber string) (*models.Products, error)
}

type GetterRepository struct {
	DB *gorm.DB
}

func NewGetterRepository(db *gorm.DB) *GetterRepository {
	return &GetterRepository{
		DB: db,
	}
}

func (repo *GetterRepository) GetAll(paging *pagination.PagingRequest) (*[]models.Products, int64, error) {
	var maxRow int64
	var products []models.Products

	err := repo.DB.Table("products").
		Where("if(? <> '', (reg_number like ? or name like ?), reg_number <> '' AND name <> '')", paging.Search, "%"+paging.Search+"%", "%"+paging.Search+"%").
		Count(&maxRow).
		Error
	if err == nil {
		err = repo.DB.Where("if(? <> '', (reg_number like ? or name like ?), name <> '')", paging.Search, "%"+paging.Search+"%", "%"+paging.Search+"%").
			Order(paging.SortAlias).
			Limit(paging.Limit).
			Offset(paging.Offset).
			Find(&products).
			Error
	}

	if err != nil {
		return nil, 0, err
	}

	return &products, maxRow, nil
}

func (repo *GetterRepository) GetById(id uint) (*models.Products, error) {
	var product models.Products
	err := repo.DB.Where("id = ?", id).
		First(&product).
		Error
	if err == nil {
		return &product, nil
	}

	return nil, custom_error.NotFound(productNotFound)
}

func (repo *GetterRepository) GetBySlug(slug string) (*models.Products, error) {
	var product models.Products
	err := repo.DB.Where("slug_name = ?", slug).
		First(&product).
		Error
	if err != nil {
		return nil, custom_error.NotFound(productNotFound)
	}

	return &product, nil
}

func (repo *GetterRepository) GetByRegNumber(regNumber string) (*models.Products, error) {
	var product models.Products
	err := repo.DB.Where("reg_number = ?", regNumber).
		First(&product).
		Error
	if err != nil {
		return nil, custom_error.NotFound(productNotFound)
	}

	return &product, nil
}
