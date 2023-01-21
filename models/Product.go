package models

type Product struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Description string `gorm:"type:text"`
	Quantity    int
	CategoryID  int
}
