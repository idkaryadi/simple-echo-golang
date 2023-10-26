package repository

import (
	"errors"
	"simple-echo-backend/config"
)

func GetProducts() ([]Products, error) {
	var products []Products
	result := config.DB.Find(&products)

	return products, result.Error
}

func GetProductById(id int) (Products, error) {
	product := Products{ID: id}
	result := config.DB.Find(&product)
	if result.RowsAffected < 1 {
		return product, errors.New("data not found")
	}
	return product, result.Error
}

func CreateProduct(newProduct Products) error {
	result := config.DB.Create(&newProduct)
	return result.Error
}

func UpdateProduct(updateProduct Products) error {
	result := config.DB.Save(&updateProduct)
	return result.Error
}

func DeleteProduct(id int) error {
	// Note: masih hard delete
	result := config.DB.Delete(&Products{}, id)
	if result.RowsAffected < 1 {
		return errors.New("data not found")
	}
	return result.Error
}
