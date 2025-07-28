package models

import "time"

type Product struct {
	ID            int64     `gorm:"column:id;primary_key" json:"id"`
	Name          string    `gorm:"column:name;not null;size:100;unique" json:"name"`
	Description   string    `gorm:"size:500" json:"description"`
	Price         float64   `gorm:"not null;type:decimal(10,2)" json:"price"`
	CostPrice     float64   `gorm:"not null;type:decimal(10,2)" json:"cost_price"`
	StockQuantity int64     `gorm:"not null" json:"stock_quantity"`
	IsActive      bool      `gorm:"not null" json:"is_active"`
	CategoryID    int64     `gorm:"not null" json:"category_id"`
	Category      Category  `gorm:"foreginKey:CategoryID" json:"category"`
	SupplierID    int64     `gorm:"not null" json:"supplier_id"`
	Supplier      Supplier  `gorm:"foreignKey:SupplierID" json:"supplier"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (p *Product) TableName() string {
	return "products"
}
