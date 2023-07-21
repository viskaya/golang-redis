package models

import (
	"aegis_test/libs/pagination"
	"time"

	"github.com/gosimple/slug"
)

type (
	Supplier struct {
		ID           uint      `gorm:"column:id;type:int;primaryKey;autoIncrement"`
		Name         string    `gorm:"column:name;type:varchar(100);not null;default:''"`
		SlugName     string    `gorm:"column:slug_name;type:varchar(100);not null;default:''"`
		PhoneNumber  string    `gorm:"column:phone;type:varchar(100);not null;default:''"`
		EmailAddress string    `gorm:"column:email_address;type:varchar(100);not null;default:''"`
		ContactName  string    `gorm:"column:contact_name;type:varchar(100);not null;default:''"`
		ContactPhone string    `gorm:"column:contact_phone;type:varchar(100);not null;default:''"`
		Address      string    `gorm:"column:address;type:text;not null;default:''"`
		CreatedAt    time.Time `gorm:"column:created_at;type:datetime;<-:create"`
		UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;<-:update"`
	}

	SupplierJSON struct {
		ID           uint       `json:"id"`
		Name         string     `json:"name" validate:"required"`
		SlugName     string     `json:"slugName"`
		PhoneNumber  string     `json:"phoneNumber" validate:"required,e164"`
		EmailAddress string     `json:"emailAddress" validate:"required,email"`
		ContactName  string     `json:"contactName" validate:"required"`
		ContactPhone string     `json:"contactPhone" validate:"required,e164"`
		Address      string     `json:"address" validate:"required"`
		CreatedAt    *time.Time `json:"createdAt"`
		UpdatedAt    *time.Time `json:"updatedAt"`
	}

	SupplierRequest struct {
		Auth     *AuthorizedRequest
		Paging   *pagination.PagingRequest
		Supplier SupplierJSON
	}
)

func (model *Supplier) ToJSON() *SupplierJSON {
	return &SupplierJSON{
		ID:           model.ID,
		Name:         model.Name,
		SlugName:     slug.Make(model.Name),
		PhoneNumber:  model.PhoneNumber,
		EmailAddress: model.EmailAddress,
		ContactName:  model.ContactName,
		ContactPhone: model.ContactPhone,
		Address:      model.Address,
	}
}

func (json *SupplierJSON) ToSupplier() Supplier {
	slug := slug.Make(json.Name)

	return Supplier{
		ID:           json.ID,
		Name:         json.Name,
		SlugName:     slug,
		PhoneNumber:  json.PhoneNumber,
		EmailAddress: json.EmailAddress,
		ContactName:  json.ContactName,
		ContactPhone: json.ContactPhone,
		Address:      json.Address,
	}
}
