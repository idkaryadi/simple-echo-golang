package core

import "time"

type Products struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	GetProducts() ([]Products, error)
	GetProductById(id int) (Products, error)
	CreateProduct(product Products) error
	UpdateProduct(product Products) error
	DeleteProduct(id int) error
}
