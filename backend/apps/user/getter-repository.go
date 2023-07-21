package user

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/pagination"
	"aegis_test/models"

	"gorm.io/gorm"
)

type GetterRepositoryInterface interface {
	GetAll(pagination *pagination.PagingRequest) (*[]models.User, int64, error)
	GetById(id uint) (*models.User, error)
	GetByUserName(userName string) (*models.User, error)
	GetBySessionId(id string) (*models.User, error)
}

type GetterRepository struct {
	DB *gorm.DB
}

func NewGetterRepository(db *gorm.DB) *GetterRepository {
	return &GetterRepository{
		DB: db,
	}
}

func (repo *GetterRepository) GetAll(pagination *pagination.PagingRequest) (*[]models.User, int64, error) {
	allUsers := []models.User{}
	var err error
	var maxRow int64
	err = repo.DB.Table("users").
		Where("if(? <> '', email LIKE ?, email <> '')", pagination.Search, "%"+pagination.Search+"%").
		Count(&maxRow).
		Error
	if err == nil {
		err = repo.DB.
			Where("if(? <> '', email LIKE ?, email <> '')", pagination.Search, "%"+pagination.Search+"%").
			Order(pagination.SortAlias).
			Limit(pagination.Limit).
			Offset(pagination.Offset).
			Find(&allUsers).
			Error
	}

	if err != nil {
		return &[]models.User{}, 0, custom_error.InternalError(err)
	}

	if len(allUsers) > 0 {
		return &allUsers, maxRow, nil
	}

	return &[]models.User{}, 0, nil
}

func (repo *GetterRepository) GetById(id uint) (*models.User, error) {
	var userModel models.User
	var err error

	err = repo.DB.First(&userModel).Error
	if err != nil {
		return nil, custom_error.InternalError(err)
	}

	return &userModel, nil
}

func (repo *GetterRepository) GetByUserName(userName string) (*models.User, error) {
	var userModel models.User
	var err error

	err = repo.DB.Where("email = ?", userName).First(&userModel).Error
	if err != nil {
		return nil, custom_error.InternalError(err)
	}

	return &userModel, nil
}

func (repo *GetterRepository) GetBySessionId(id string) (*models.User, error) {
	var userModel models.User

	err := repo.DB.Where("session_id = ?", id).Find(&userModel).Error
	if err != nil {
		return nil, custom_error.InternalError(err)
	}

	return &userModel, nil
}
