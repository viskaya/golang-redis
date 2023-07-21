package user

import (
	"aegis_test/libs"
	"aegis_test/libs/custom_error"
	"aegis_test/libs/db"
	"aegis_test/models"
)

type LoginServiceInterface interface {
	LoginByUserName(userName string, password string) (*models.UserAccessAuthorizedJSON, error)
}

type LoginService struct {
	Cache            db.CacheDB
	Repository       LoginRepositoryInterface
	SessionGenerator libs.UserSessionGeneratorInterface
}

func NewLoginService(db *db.DBFactory) *LoginService {
	repo := NewLoginRepository(db)
	sessionGen := libs.NewUserSessionGenerator()

	return &LoginService{
		Cache:            db.Cache,
		Repository:       repo,
		SessionGenerator: sessionGen,
	}
}

func (service *LoginService) LoginByUserName(userName string, password string) (*models.UserAccessAuthorizedJSON, error) {
	user, err := service.Repository.GetByUserName(userName)

	if err != nil {
		return nil, custom_error.InternalError(err)
	} else {
		if password != user.Password {
			return nil, custom_error.BadRequest("invalid user name or password")
		} else if user.Status != models.UserActive {
			if user.Status == models.UserSuspended {
				return nil, custom_error.BadRequest("your account is suspended. please contact administrator to activate")
			} else if user.Status == models.UserInactive {
				return nil, custom_error.BadRequest("your account is inactive. please contact administrator to activate")
			}
		}
	}

	service.Cache.RemoveCacheForContext("auth")
	service.Cache.RemoveCacheForContext(cacheContextName)

	sessionId, sessionExpire := service.SessionGenerator.SessionGenerate(user.ID)

	user.SessionExpireAt = *sessionExpire
	user.SessionID = *sessionId
	service.Repository.UpdateAfterLogin(&user)

	var profiles []string
	for _, profile := range user.UserProfiles {
		profiles = append(profiles, profile.Profile.Role)
	}

	token := libs.CreateJWTToken(*sessionId, *sessionExpire, profiles)

	return user.ToAccessAuthorizedJSON(token), nil
}
