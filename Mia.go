package entity

import "time"

// Lets make the Roster and LiveGame the primary key here
type Mia struct {
	IdRoster   uint `gorm:"column:idroster;primary_key"`
	IdLiveGame uint `gorm:"column:idlivegame;primary_key"`

	// Control information
	ID        uint      `gorm:"column:idmia;AUTO_INCREMENT"`
	IsActive  *bool      `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time `gorm:"-"`
}
