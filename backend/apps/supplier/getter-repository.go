package supplier

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/pagination"
	"aegis_test/models"

	"gorm.io/gorm"
)

const (
	fieldCompanyId string = "company_id = ?"
)

type GetterRepositoryInterface interface {
	GetAll(paging *pagination.PagingRequest) (*[]models.Supplier, int64, error)
	GetById(id uint) (*models.Supplier, error)
	GetBySlug(slug string) (*models.Supplier, error)
}

type GetterRepository struct {
	DB *gorm.DB
}

func NewGetterRepository(db *gorm.DB) *GetterRepository {
	return &GetterRepository{
		DB: db,
	}
}

func (repo *GetterRepository) GetAll(paging *pagination.PagingRequest) (*[]models.Supplier, int64, error) {
	var maxRow int64
	var suppliers []models.Supplier

	err := repo.DB.Table("suppliers").
		Where("if(? <> '', name like ?, name <> '')", paging.Search, "%"+paging.Search+"%").
		Count(&maxRow).
		Error
	if err == nil {
		err = repo.DB.Where("if(? <> '', name like ?, name <> '')", paging.Search, "%"+paging.Search+"%").
			Order(paging.SortAlias).
			Limit(paging.Limit).
			Offset(paging.Offset).
			Find(&suppliers).
			Error
	}

	if err != nil {
		return nil, 0, err
	}

	return &suppliers, maxRow, nil
}

func (repo *GetterRepository) GetById(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	err := repo.DB.Where("id = ?", id).
		First(&supplier).
		Error
	if err != nil {
		return nil, custom_error.NotFound(supplierNotFound)
	}

	return &supplier, nil
}

func (repo *GetterRepository) GetBySlug(slug string) (*models.Supplier, error) {
	var supplier models.Supplier
	err := repo.DB.Where("slug_name = ?", slug).
		First(&supplier).
		Error
	if err != nil {
		return nil, custom_error.NotFound(supplierNotFound)
	}

	return &supplier, nil
}
