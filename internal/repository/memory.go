package repository

import (
	"context"
	"shopping-backend/internal/models"
	"sync"
)

// In Memory implementation for quick testing
type MemoryProductRepo struct {
	mu       sync.RWMutex
	products map[string]*models.Product
}

func NewMemoryProductRepo() *MemoryProductRepo {
	return &MemoryProductRepo{
		products: make(map[string]*models.Product),
	}
}

func (mpr *MemoryProductRepo) AddProduct(ctx context.Context, product *models.Product) error {
	mpr.mu.Lock()
	mpr.products[product.ID] = product
	mpr.mu.Unlock()
	return nil
}

func (mpr *MemoryProductRepo) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	mpr.mu.RLock()
	defer mpr.mu.RUnlock()
	product, exists := mpr.products[id]
	if !exists {
		return nil, models.ErrProductNotFound
	}
	return product, nil
}
func (mpr *MemoryProductRepo) DeleteProductByID(ctx context.Context, id string) error {
	mpr.mu.Lock()
	defer mpr.mu.Unlock()
	if _, exists := mpr.products[id]; !exists {
		return models.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}

func (mpr *MemoryProductRepo) UpdateProductByID(ctx context.Context, id string, delta int) error {
	mpr.mu.Lock()
	defer mpr.mu.Unlock()
	product, exists := mpr.products[id]
	if !exists {
		return models.ErrProductNotFound
	}
	if product.Stock+delta < 0 {
		return models.ErrInsufficientStock
	}
	product.Stock += delta
	return nil
}

type MemoryCartRepo struct {
	mu    sync.RWMutex
	carts map[string]*models.Cart
}

func NewMemoryCartRepo() *MemoryCartRepo {
	return &MemoryCartRepo{
		carts: make(map[string]*models.Cart),
	}
}

func (mcr *MemoryCartRepo) GetCartByUserID(ctx context.Context, userID string) ([]*models.Cart, error) {
	mcr.mu.RLock()
	defer mcr.mu.RUnlock()
	carts := make([]*models.Cart, 0)
	for _, cart := range mcr.carts {
		if cart.UserID == userID {
			cartCopy := &cart
			carts = append(carts, *cartCopy)
		}
	}
	return carts, nil
}

func (mcr *MemoryCartRepo) GetCartByID(ctx context.Context, cartID string) (*models.Cart, error) {
	mcr.mu.RLock()
	defer mcr.mu.RUnlock()
	cart, exists := mcr.carts[cartID]
	if !exists {
		return nil, models.ErrCartNotFound
	}
	cartCopy := *cart
	return &cartCopy, nil
}

func (mcr *MemoryCartRepo) SaveCart(ctx context.Context, cart *models.Cart) error {
	mcr.mu.Lock()
	defer mcr.mu.Unlock()
	copyCart := *cart
	mcr.carts[cart.ID] = &copyCart

	return nil
}
func (mcr *MemoryCartRepo) Clear(ctx context.Context, cartID string) error {
	mcr.mu.Lock()
	defer mcr.mu.Unlock()
	cart, exists := mcr.carts[cartID]
	if !exists {
		return models.ErrCartNotFound
	}
	cart.CartItems = make(map[string]models.CartItem)
	return nil
}
func (mcr *MemoryCartRepo) DeleteCartByID(ctx context.Context, cartID string) error {
	mcr.mu.Lock()
	defer mcr.mu.Unlock()
	if _, exists := mcr.carts[cartID]; !exists {
		return models.ErrCartNotFound
	}
	delete(mcr.carts, cartID)
	return nil
}
func (mcr *MemoryCartRepo) DeleteItemFromCart(ctx context.Context, cartID string, cartItemID string) error {
	mcr.mu.Lock()
	defer mcr.mu.Unlock()
	cart, exists := mcr.carts[cartID]
	if !exists {
		return models.ErrCartNotFound
	}
	delete(cart.CartItems, cartItemID)
	return nil
}
