package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code       string
	Price      uint
	TestField1 int64 `gorm:"column:testField1"` // set column name to `testField1`
	TestField2 int64
}

// Set Product's table name to be `prodotti`
func (Product) TableName() string {
	return "prodotti"
}
