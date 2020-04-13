package entity

import "time"

type Mia struct {
	IdRoster   uint `gorm:"column:idroster"`
	IdLiveGame uint `gorm:"column:idlivegame"`

	// Controll information
	ID        uint      `gorm:"column:idmia;primary_key"`
	IsActive  bool      `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}
