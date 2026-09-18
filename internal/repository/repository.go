package repository

import (
	"context"
	"shopping-backend/internal/models"
)

type ProductRepository interface {
	AddProduct(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	DeleteProductByID(ctx context.Context, id string) error
	UpdateProductByID(ctx context.Context, id string, delta int) error
}
type CartRepository interface {
	GetCartByUserID(ctx context.Context, userID string) (*models.Cart, error)
	GetCartByID(ctx context.Context, cartID string) (*models.Cart, error)
	SaveCart(ctx context.Context, cart *models.Cart) error
	Clear(ctx context.Context, cartID string) error
	DeleteCartByID(ctx context.Context, cartID string) error
	DeleteItemFromCart(ctx context.Context, cartItemID string) error
}
type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	GetOrderByUserID(ctx context.Context, userID string) (*models.Order, error)
}
type UserRepository interface {
	AddUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}
