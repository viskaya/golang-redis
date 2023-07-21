package user

import (
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/libs/pagination"
	"aegis_test/models"
)

const (
	cacheContextName string = "user"
)

type GetterServiceInterface interface {
	GetAll(request *models.UserRequest) *pagination.PagingResponseJSON
	GetById(request *models.UserRequest) (*models.UserJSON, error)
	GetByUserName(request *models.UserRequest) (*models.UserJSON, error)
	AuthBySessionId(id string) (*models.AuthorizedUser, error)
}

type GetterService struct {
	Cache      db.CacheDB
	Repository GetterRepositoryInterface
}

func NewGetterService(db *db.DBFactory) *GetterService {
	return &GetterService{
		Cache:      db.Cache,
		Repository: NewGetterRepository(db.DB),
	}
}

func (service *GetterService) GetAll(request *models.UserRequest) *pagination.PagingResponseJSON {
	paging := request.Paging

	var page *pagination.PagingResponseJSON

	cacheKey := service.Cache.CacheKeyBuilder(cacheContextName, "GetterService", "GetAll", paging.PagingToString())
	cacheData, err := service.Cache.GetCache(cacheKey, &pagination.PagingResponseJSON{})
	if err != nil {
		var users []interface{}
		userModels, maxRow, _ := service.Repository.GetAll(paging)
		if userModels != nil {
			for _, user := range *userModels {
				users = append(users, *user.ToJSON())
			}
		}
		page := pagination.NewPagingResponse(paging, users, maxRow)
		if page.MaxRow > 0 {
			service.Cache.SetCache(cacheKey, page.GetPagingResponse())
		}
	} else {
		page = cacheData.(*pagination.PagingResponseJSON)
	}

	return page
}

func (service *GetterService) GetById(request *models.UserRequest) (*models.UserJSON, error) {
	requestId := request.User.ID
	user, err := service.Repository.GetById(requestId)

	if err != nil {
		return nil, custom_error.InternalError(err)
	}

	return user.ToJSON(), nil
}

func (service *GetterService) GetByUserName(request *models.UserRequest) (*models.UserJSON, error) {
	requserUserName := request.User.UserName
	user, err := service.Repository.GetByUserName(requserUserName)

	if err != nil {
		return nil, custom_error.InternalError(err)
	}

	return user.ToJSON(), nil
}

func (service *GetterService) AuthBySessionId(id string) (*models.AuthorizedUser, error) {
	user, err := service.Repository.GetBySessionId(id)

	if err == nil && user.ID > 0 {
		return user.ToAuthorizedUser(), nil
	}

	return nil, custom_error.InternalError(err)
}
