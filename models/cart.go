package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        int64          `gorm:"column:id;primaryKey" json:"id"`
	UserID    int64          `gorm:"not null;unique" json:"user_id"`
	CartItems []CartItem     `json:"cart_items,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (c *Cart) TableName() string {
	return "carts"
}

type CartItem struct {
	ID               int64          `gorm:"column:id;primaryKey" json:"id"`
	CartID           int64          `gorm:"not null" json:"cart_id"`
	Cart             Cart           `gorm:"foreignKey:CartID" json:"cart"`
	ProductID        int64          `gorm:"not null" json:"product_id"`
	Product          Product        `gorm:"foreignKey:ProductID" json:"product"`
	Quantity         int            `gorm:"not null;default:1" json:"quantity"`
	PriceAtAddToCart float64        `gorm:"type:decimal(10,2);not null" json:"price_at_add_to_cart"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdatetime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (ci *CartItem) TableName() string {
	return "cart_items"
}

// For manage shopping cart operations
type CartController struct {
	DB *gorm.DB
}

// For make new instance from CartController
func NewCartController(db *gorm.DB) *CartController {
	return &CartController{DB: db}
}
