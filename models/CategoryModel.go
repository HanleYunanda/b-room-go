package models

type Category struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
