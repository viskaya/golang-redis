package models

import (
	"aegis_test/libs/pagination"
	"time"

	"github.com/gosimple/slug"
)

type (
	Products struct {
		ID             uint       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
		RegNumber      string     `gorm:"column:reg_number;type:varchar(100);not null;default:''"`
		Name           string     `gorm:"column:name;type:varchar(100);not null;default:''"`
		SlugName       string     `gorm:"column:slug_name;type:varchar(100);not null;default:''"`
		Photo          string     `gorm:"column:photo;type:varchar(100);not null;default:''"`
		SupplierID     uint       `gorm:"column:supplier_id;type:int;not null;default:0"`
		UnitID         uint       `gorm:"column:unit_id;type:int;not null;default:0"`
		Quantity       float64    `gorm:"column:quantity;type:double(14,2);not null;default:0.0"`
		BookedQuantity float64    `gorm:"column:booked_quantity;type:double(14,2);not null;default:0.0"`
		BasicPrice     float64    `gorm:"column:basic_price;type:double(14,2);not null;default:0.0"`
		SellingPrice   float64    `gorm:"column:selling_price;type:double(14,2);not null;default:0.0"`
		CreatedAt      *time.Time `gorm:"column:created_at;type:datetime;<-:create"`
		UpdatedAt      *time.Time `gorm:"column:updated_at;type:datetime;<-:update"`
	}

	ProductJSON struct {
		ID             uint       `json:"id"`
		RegNumber      string     `json:"regNumber" validate:"required"`
		Name           string     `json:"name" validate:"required"`
		SlugName       string     `json:"slugName"`
		Photo          string     `json:"photo"`
		SupplierID     uint       `json:"supplierId" validate:"required,gt=0"`
		UnitID         uint       `json:"unitId" validate:"required,gt=0"`
		Quantity       float64    `sql:"type:double(14,2)" json:"quantity" validate:"required,gt=0.0"`
		BookedQuantity float64    `sql:"type:double(14,2)" json:"bookedQuantity"`
		BasicPrice     float64    `sql:"type:double(14,2)" json:"basicPrice" validate:"required,gt=0.0"`
		SellingPrice   float64    `sql:"type:double(14,2)" json:"sellingPrice" validate:"required,gt=0.0"`
		CreatedAt      *time.Time `json:"createdAt;"`
		UpdatedAt      *time.Time `json:"updatedAt;"`
	}

	ProductRequest struct {
		Auth    *AuthorizedRequest
		Paging  *pagination.PagingRequest
		Product ProductJSON
	}
)

func (model *Products) ToJSON() *ProductJSON {
	return &ProductJSON{
		ID:           model.ID,
		RegNumber:    model.RegNumber,
		Name:         model.Name,
		SlugName:     model.SlugName,
		Photo:        model.Photo,
		SupplierID:   model.SupplierID,
		UnitID:       model.UnitID,
		Quantity:     model.Quantity,
		BasicPrice:   model.BasicPrice,
		SellingPrice: model.SellingPrice,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

func (json *ProductJSON) ToProduct() *Products {
	slug := slug.Make(json.Name)

	return &Products{
		ID:           json.ID,
		RegNumber:    json.RegNumber,
		Name:         json.Name,
		SlugName:     slug,
		Photo:        json.Photo,
		SupplierID:   json.SupplierID,
		UnitID:       json.UnitID,
		Quantity:     json.Quantity,
		BasicPrice:   json.BasicPrice,
		SellingPrice: json.SellingPrice,
	}
}
