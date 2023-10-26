package products

import (
	"time"
)

type Products struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// DeletedAt   *time.Time `gorm:"default:NULL"`
}
