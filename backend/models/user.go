package models

import (
	"aegis_test/libs/pagination"
	"time"
)

const (
	UserActive    string = "Active"
	UserInactive  string = "Inactive"
	UserSuspended string = "Suspended"
)

type (
	UserInterface interface {
		ToUserJSON() *UserJSON
		ToAuthorizedUser() *AuthorizedUser
		ToAccessAuthorizedJSON(accessToken string) *UserAccessAuthorizedJSON
	}

	User struct {
		ID              uint       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
		UserName        string     `gorm:"column:email;type:varchar(100)"`
		Password        string     `gorm:"column:password;type:varchar(100)"`
		Status          string     `gorm:"column:login_status;type:enum('Active','Inactive','Suspend')"`
		SessionID       string     `gorm:"column:session_id;type:varchar(255)"`
		SessionExpireAt time.Time  `gorm:"column:session_expire_at;type:datetime"`
		SuspendedAt     *time.Time `gorm:"column:suspend_at;type:datetime"`
		LastLoginAt     time.Time  `gorm:"column:last_login_at;type:datetime"`
		ActivationToken string     `gorm:"column:activation_token;type:varchar(255)"`
		CreatedAt       time.Time  `gorm:"column:created_at;type:datetime;<-:create"`
		UpdatedAt       time.Time  `gorm:"column:updated_at;type:datetime;<-:update"`
		UserProfiles    []UserProfile
	}

	UserProfile struct {
		ID        uint      `gorm:"column:id;type:int"`
		UserID    uint      `gorm:"column:user_id;type:int"`
		ProfileID uint      `gorm:"column:profile_id;type:int"`
		CreatedAt time.Time `gorm:"column:created_at;type:datetime;<-:create"`
		UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;<-:update"`
		Profile   Profile
	}

	Profile struct {
		ID        uint      `gorm:"column:id;type:int"`
		Role      string    `gorm:"column:role;type:varchar(100)"`
		Name      string    `gorm:"column:name;type:varchar(100)"`
		CreatedAt time.Time `gorm:"column:created_at;type:datetime;<-:create"`
		UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;<-:update"`
	}

	UserJSON struct {
		ID          uint       `json:"id"`
		UserName    string     `json:"userName"`
		Status      string     `json:"status"`
		SuspendedAt *time.Time `json:"suspendedAt"`
		LastLoginAt time.Time  `json:"lastLoginAt"`
	}

	UserSaveJSON struct {
		ID       uint   `json:"id"`
		UserName string `json:"userName" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	UserAccessAuthorizedJSON struct {
		ID             uint       `json:"id"`
		UserName       string     `json:"userName"`
		Status         string     `json:"status"`
		SuspendedAt    *time.Time `json:"suspendedAt"`
		LastLoginAt    time.Time  `json:"lastLoginAt"`
		AccessToken    string     `json:"accessToken,omitempty"`
		AccessExpireAt time.Time  `json:"accessExpireAt"`
	}

	AuthorizedUser struct {
		ID       uint
		UserName string
		Status   string
		Roles    []string
		Profile  []string
	}

	UserRequest struct {
		Auth   *AuthorizedRequest
		Paging *pagination.PagingRequest
		User   UserJSON
	}

	UserSaveRequest struct {
		Auth *AuthorizedRequest
		User UserSaveJSON
	}
)

func (model *User) ToJSON() *UserJSON {
	return &UserJSON{
		ID:          model.ID,
		UserName:    model.UserName,
		Status:      model.Status,
		SuspendedAt: &model.SessionExpireAt,
		LastLoginAt: model.LastLoginAt,
	}
}

func (json *UserJSON) ToUser() *User {
	return &User{
		ID:       json.ID,
		UserName: json.UserName,
	}
}

func (json *UserSaveJSON) ToUser() *User {
	return &User{
		ID:       json.ID,
		UserName: json.UserName,
		Password: json.Password,
	}
}

func (model *User) ToAccessAuthorizedJSON(accessToken string) *UserAccessAuthorizedJSON {
	return &UserAccessAuthorizedJSON{
		ID:             model.ID,
		UserName:       model.UserName,
		Status:         model.Status,
		SuspendedAt:    model.SuspendedAt,
		LastLoginAt:    model.LastLoginAt,
		AccessToken:    accessToken,
		AccessExpireAt: model.SessionExpireAt,
	}
}

func (model *User) ToAuthorizedUser() *AuthorizedUser {
	var roles []string
	var profiles []string

	if model.UserProfiles != nil {
		for _, role := range model.UserProfiles {
			roles = append(roles, role.Profile.Role)
		}
		for _, profile := range model.UserProfiles {
			profiles = append(profiles, profile.Profile.Name)
		}
	}

	return &AuthorizedUser{
		ID:       model.ID,
		UserName: model.UserName,
		Status:   model.Status,
		Roles:    roles,
		Profile:  profiles,
	}
}
