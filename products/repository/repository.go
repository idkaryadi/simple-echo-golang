package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
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

func (repo Repository) GetProducts() ([]Products, error) {
	var products []Products
	result := repo.db.Find(&products)

	return products, result.Error
}

func (repo Repository) GetProductById(id int) (Products, error) {
	product := Products{ID: id}
	result := repo.db.Find(&product)
	if result.RowsAffected < 1 {
		return product, errors.New("data not found")
	}
	return product, result.Error
}

func (repo Repository) CreateProduct(newProduct Products) error {
	result := repo.db.Create(&newProduct)
	return result.Error
}

func (repo Repository) UpdateProduct(updateProduct Products) error {
	result := repo.db.Save(&updateProduct)
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
