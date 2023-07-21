package models

import (
	"aegis_test/libs"
	"time"
)

type (
	Cashiers struct {
		ID           uint       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
		ProductID    uint       `gorm:"column:product_id;type:int;not null;default:0;index:cashier_productid_idx"`
		TrxNumber    string     `gorm:"column:trx_number;type:varchar(100);not null;default:'';index:cashier_trxnumber_idx"`
		TrxDate      time.Time  `gorm:"column:trx_date;type:datetime;not null;default:'';index:cashier_trxdate_idx;<-:create"`
		TotalPayment float64    `gorm:"column:total_payment;type:double(14,2);not null;default:0.0"`
		Quantity     float64    `gorm:"column:quantity;type:double(14,2);not null;default:0.0"`
		CreatedAt    *time.Time `gorm:"column:created_at;type:datetime;<-:create"`
		UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime;<-:update"`
	}

	CashierJSON struct {
		ID           uint    `json:"id"`
		ProductID    uint    `json:"productId" validate:"required,gt=0"`
		TrxNumber    string  `json:"trxNumber"`
		TrxDate      string  `json:"trxDate"`
		TotalPayment float64 `sql:"type:double(14,2)" json:"totalPayment" validate:"required,gt=0"`
		Quantity     float64 `sql:"type:double(14,2)" json:"quantity" validate:"required,gt=0"`
	}

	CashierRequest struct {
		Auth    *AuthorizedRequest
		Cashier *CashierJSON
	}
)

func (model *Cashiers) ToJSON() *CashierJSON {
	return &CashierJSON{
		ID:           model.ID,
		ProductID:    model.ProductID,
		TrxNumber:    model.TrxNumber,
		TrxDate:      *libs.FormatDateTime(model.TrxDate),
		TotalPayment: model.TotalPayment,
		Quantity:     model.Quantity,
	}
}

func (json *CashierJSON) ToCashier() *Cashiers {
	return &Cashiers{
		ID:           json.ID,
		ProductID:    json.ProductID,
		TrxNumber:    json.TrxNumber,
		TotalPayment: json.TotalPayment,
		Quantity:     json.Quantity,
	}
}
