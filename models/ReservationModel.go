package models

type Reservation struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	UserId   int64  `gorm:"type:int" json:"user_id"`
	User     User   `json:"user"`
	RoomId   int64  `gorm:"type:int" json:"room_id"`
	Room     Room   `json:"room"`
	CheckIn  string `gorm:"type:timestamp" json:"check_in"`
	CheckOut string `gorm:"type:timestamp" json:"check_out"`
}
