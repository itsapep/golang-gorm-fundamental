package entity

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string      `gorm:"column:name;not null"`
	Customers   []*Customer `gorm:"many2many:customer_products"`
}

func (p Product) TableName() string {
	return "mst_product"
}

func (c *Product) ToString() string {
	customer, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return ""
	}
	return string(customer)
}
