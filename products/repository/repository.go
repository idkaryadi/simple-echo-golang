package repository

import (
	"errors"
	"simple-echo-backend/products/core"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	// TODO: auto migrate
	db.AutoMigrate(&Products{})

	return Repository{db: db}
}

type Products struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// DeletedAt   *time.Time `gorm:"default:NULL"`
}

func (p Products) toCoreProducts() core.Products {
	return core.Products{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func newProductDoc(core core.Products) Products {
	return Products{
		ID:          core.ID,
		Name:        core.Name,
		Description: core.Description,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
	}
}

func (repo Repository) GetProducts() ([]core.Products, error) {
	var products []Products
	result := repo.db.Find(&products)

	productsCore := make([]core.Products, len(products))
	for i, v := range products {
		productsCore[i] = v.toCoreProducts()
	}

	return productsCore, result.Error
}

func (repo Repository) GetProductById(id int) (core.Products, error) {
	product := Products{ID: id}
	result := repo.db.Find(&product)
	if result.RowsAffected < 1 {
		return product.toCoreProducts(), errors.New("data not found")
	}
	return product.toCoreProducts(), result.Error
}

func (repo Repository) CreateProduct(newProduct core.Products) error {
	product := newProductDoc(newProduct)
	result := repo.db.Create(&product)
	return result.Error
}

func (repo Repository) UpdateProduct(updateProduct core.Products) error {
	product := newProductDoc(updateProduct)
	result := repo.db.Save(&product)
	return result.Error
}

func (repo Repository) DeleteProduct(id int) error {
	// Note: masih hard delete
	result := repo.db.Delete(&Products{}, id)
	if result.RowsAffected < 1 {
		return errors.New("data not found")
	}
	return result.Error
}
