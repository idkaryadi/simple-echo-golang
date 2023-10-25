package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Products struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// DeletedAt   *time.Time `gorm:"default:NULL"`
}

var (
	DB  *gorm.DB
	err error
)

func connectDB() {
	dsn := "root:simple-password@tcp(mysqldb:3306)/simple-project?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Err::", err)
	}

	// TODO: auto migrate
	DB.AutoMigrate(&Products{})
}
