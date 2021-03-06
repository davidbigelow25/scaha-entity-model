package entity

import "time"

type Player struct {
	IdPerson     uint `gorm:"column:idperson"`
	IsGoalie     uint `gorm:"column:isgoalie"`
	IsBcVerified uint `gorm:"column:bcverified"`

	// Controll information
	ID        uint      `gorm:"column:idplayer;primary_key;AUTO_INCREMENT"`
	IsActive  *bool     `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}
