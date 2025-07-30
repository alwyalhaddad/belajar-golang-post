package models

import "time"

type StockAdjustment struct {
	ID             int64          `gorm:"column:id;primary_key" json:"id"`
	ProductID      int64          `gorm:"not null" json:"product_id"`
	Product        Product        `gorm:"foreignKey:ProductID" json:"product"`
	UserID         int64          `gorm:"not null" json:"user_id"`
	User           User           `gorm:"foreignKey:UserID" json:"user"`
	AdjustmentType AdjustmentType `gorm:"embedded"`
	QuantityChange int64          `gorm:"not null" json:"quantity_change"`
	Reason         string         `gorm:"not null" json:"reason"`
	AdjustmentDate time.Time      `gorm:"type:date" json:"adjustment_date"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
}

func (sa *StockAdjustment) TableName() string {
	return "stock_adjustments"
}

type AdjustmentType struct {
	Add                  string `gorm:"column:add"`
	Remove               string `gorm:"column:remove"`
	Damaged              string `gorm:"column:damaged"`
	ReceivedFromSupplier string `gorm:"column:received_from_supplier"`
}
