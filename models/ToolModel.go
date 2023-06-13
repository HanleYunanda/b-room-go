package models

type Tool struct {
	Id            int    `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	ReservationId int    `json:"reservation_id"`
}
