package models

type Reservation struct {
	Id              int64  `gorm:"primaryKey" json:"id"`
	UserId          int64  `gorm:"type:int" json:"user_id"`
	User            User   `json:"user"`
	RoomId          int64  `gorm:"type:int" json:"room_id"`
	Room            Room   `json:"room"`
	ReservationDate string `gorm:"type:date" json:"reservation_date"`
	CheckIn         string `gorm:"type:time" json:"check_in"`
	CheckOut        string `gorm:"type:time" json:"check_out"`
	Tools           []Tool `json:"tools"`
}
