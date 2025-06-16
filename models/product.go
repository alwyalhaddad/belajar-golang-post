package models

import "time"

type Product struct {
	ProductID     int64     `gorm:"column:product_id;primary_key"`
	Sku           int64     `gorm:"column:stock_keeping_unit;<-:create"`
	Name          string    `gorm:"column:name"`
	Description   string    `gorm:"column:description"`
	Price         int64     `gorm:"column:price"`
	CostPrice     int64     `gorm:"column:cost_price"`
	StockQuantity int64     `gorm:"column:stock_quantity"`
	IsActive      bool      `gorm:"column:is_active"`
	CategoryID    Category  `gorm:"foreignKey:CategoryID;references:category_id"`
	SupplierID    Supplier  `gorm:"foreignKey:SupplierID;references:supplier_id"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
