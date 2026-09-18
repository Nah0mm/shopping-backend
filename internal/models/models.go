package models

import (
	"errors"
	"time"
)

type User struct {
	ID           string  `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Address      string  `json:"address"`
	Balance      float64 `json:"balance"`
	OrderHistory []Order `json:"order_history"`
	Carts        []Cart  `json:"carts"`
}

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	Stock     int       `json:"stock"`
}

type Cart struct {
	ID        string              `json:"cart_id"`
	UserID    string              `json:"user_id"`
	CartItems map[string]CartItem `json:"items"`
}

type CartItem struct {
	ProductID    string  `json:"product_id"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
}

type Order struct {
	ID          string      `json:"order_id"`
	UserID      string      `json:"user_id"`
	CreatedAt   time.Time   `json:"created_at"`
	TotalAmount float64     `json:"total_amount"`
	OrderItems  []OrderItem `json:"order_items"`
	Status      OrderStatus `json:"status"`
}

type OrderItem struct {
	ProductID    string  `json:"product_id"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
}

var (
	ErrInsufficientStock = errors.New("Insufficient amount of product")
	ErrCartIsEmpty       = errors.New("Cart is empty")
	ErrProductNotFound   = errors.New("Product not found")
	ErrInsufficientFunds = errors.New("Insufficient funds")
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusCompleted OrderStatus = "COMPLETED"
	StatusCancelled OrderStatus = "CANCELLED"
)
