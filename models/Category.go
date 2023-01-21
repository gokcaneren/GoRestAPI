package models

type Category struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Description string `gorm:"type:text"`
	Products    []Product
}
