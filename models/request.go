package models

// Request to login_controller.go
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Request to register_controller.go
type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

// Request to change_password_controller.go
type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password" binding:"required"`
	NewPassword        string `json:"new_password" binding:"required,min=9"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}

// Request to forgot_password_controller.go
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// Request to product.go
type CreateProductRequest struct {
	Name          string  `json:"name" binding:"omitempty,min=3,max=100"`
	Description   string  `json:"description" binding:"max=500"`
	Price         float64 `json:"price" binding:"required,min=0"`
	CostPrice     int64   `json:"cost_price" binding:"required,min=0"`
	StockQuantity int64   `json:"stock_quantity" binding:"required,min=0"`
	IsActive      bool    `json:"is_active"`
	CategoryID    int64   `json:"category_id" binding:"required"`
	SupplierID    int64   `json:"supplier_id" binding:"required"`
}

type UpdateProductRequest struct {
	Name          *string  `json:"name" binding:"omitempty,min=3,max=100"`
	Description   *string  `json:"description" binding:"omitempty,max=500"`
	Price         *float64 `json:"price" binding:"omitempty,min=0"`
	CostPrice     *int64   `json:"cost_price" binding:"omitempty,min=0"`
	StockQuantity *int64   `json:"stock_quantity" binding:"omitempty,min=0"`
	IsActive      *bool    `json:"is_active" binding:"omitempty"`
	CategoryID    *uint    `json:"category_id" binding:"omitempty"`
	SupplierID    *uint    `json:"supplier_id" binding:"omitempty"`
}
