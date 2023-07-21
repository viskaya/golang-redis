package models

import (
	"aegis_test/libs/pagination"
	"time"

	"github.com/gosimple/slug"
)

type (
	Unit struct {
		ID        uint      `gorm:"column:id;type:int;primaryKey;autoIncrement"`
		Name      string    `gorm:"column:name;type:varchar(100);not null;default:''"`
		SlugName  string    `gorm:"column:slug_name;type:varchar(100);not null;default:''"`
		CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:datetime;<-:create"`
		UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:datetime;<-:update"`
	}

	UnitJSON struct {
		ID       uint   `json:"id"`
		Name     string `json:"name" validate:"required"`
		SlugName string `json:"slugName"`
	}

	UnitRequest struct {
		Auth   *AuthorizedRequest
		Paging *pagination.PagingRequest
		Unit   UnitJSON
	}
)

func (unit *Unit) ToJSON() *UnitJSON {
	return &UnitJSON{
		ID:       unit.ID,
		Name:     unit.Name,
		SlugName: slug.Make(unit.Name),
	}
}

func (json *UnitJSON) ToUnit() Unit {

	return Unit{
		ID:       json.ID,
		Name:     json.Name,
		SlugName: slug.Make(json.Name),
	}
}
