package models

type Room struct {
	Id         int      `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"type:varchar(255)" json:"name"`
	CategoryId int      `gorm:"type:int" json:"category_id"`
	Category   Category `json:"category"`
}
