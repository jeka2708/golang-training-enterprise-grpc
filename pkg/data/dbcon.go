package data

import "gorm.io/gorm"

type DataEnterprise struct {
	db *gorm.DB
}

func NewDataEnterprise(db *gorm.DB) *DataEnterprise {
	return &DataEnterprise{db: db}
}
