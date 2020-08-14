package entity

import "time"

type Penalty struct {
	IdLiveGame uint `gorm:"column:idlivegame"`
	IdTeam     uint `gorm:"column:idteam"`
	IdRoster   uint `gorm:"column:idroster"`

	Period        uint
	PenaltyType   string `gorm:"column:penaltytype"`
	Minutes       string
	TimeOfPenalty string `gorm:"column:timeofpenalty"`

	// Controll information
	ID        uint      `gorm:"column:idpenalty;primary_key;AUTO_INCREMENT"`
	IsActive  *bool     `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time `gorm:"-"`
}
