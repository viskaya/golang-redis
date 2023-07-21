package unit

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/pagination"
	"aegis_test/models"
	"fmt"

	"gorm.io/gorm"
)

type GetterRepositoryInterface interface {
	GetAll(paging *pagination.PagingRequest) (*[]models.Unit, int64, error)
	GetById(id uint) (*models.Unit, error)
	GetBySlug(slug string) (*models.Unit, error)
}

type GetterRepository struct {
	DB *gorm.DB
}

func NewGetterRepository(db *gorm.DB) *GetterRepository {
	return &GetterRepository{
		DB: db,
	}
}

func (repo *GetterRepository) GetAll(paging *pagination.PagingRequest) (*[]models.Unit, int64, error) {
	var maxRow int64
	var units []models.Unit

	err := repo.DB.Table("units").
		Where("if(? <> '', name like ?, name <> '')", paging.Search, fmt.Sprintf("%s%s%s", "%", paging.Search, "%")).
		Count(&maxRow).
		Error
	if err == nil {
		err = repo.DB.Where("if(? <> '', name like ?, name <> '')", paging.Search, fmt.Sprintf("%s%s%s", "%", paging.Search, "%")).
			Order(paging.SortAlias).
			Limit(paging.Limit).
			Offset(paging.Offset).
			Find(&units).
			Error
	}

	if err != nil {
		return nil, 0, err
	}

	return &units, maxRow, nil
}

func (repo *GetterRepository) GetById(id uint) (*models.Unit, error) {
	var unit models.Unit
	err := repo.DB.Where("id = ?", id).
		First(&unit).
		Error
	if err != nil {
		return nil, custom_error.NotFound("unit could not found")
	}

	return &unit, nil
}

func (repo *GetterRepository) GetBySlug(slug string) (*models.Unit, error) {
	var unit models.Unit
	err := repo.DB.Where("slug_name = ?", slug).
		First(&unit).
		Error
	if err != nil {
		return nil, custom_error.NotFound("unit could not found")
	}

	return &unit, nil
}
