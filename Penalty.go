package entity

import "time"

type Penalty struct {

	IdLiveGame uint `gorm:"column:idlivegame"`
	IdTeam uint `gorm:"column:idteam"`
	IdRoster uint `gorm:"column:idroster"`

	IdRosterGoal uint `gorm:"column:idroster_goal"`
	IdRosterA1 uint `gorm:"column:idroster_a1"`
	IdRosterA2 uint `gorm:"column:idroster_a2"`

	Period   uint
	PenaltyType string `gorm:"column:penaltytype"`
	Minutes string
	TimeOfPenalty string `gorm:"column:timeofpenalty"`

	// Controll information
	ID uint `gorm:"column:idpenalty;primary_key"`
	IsActive bool    `gorm:"column:isactive"`
	UpdatedAt  time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt  time.Time

}
