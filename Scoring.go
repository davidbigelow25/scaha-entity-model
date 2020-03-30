package entity

import (
	"time"
)

type Scoring struct {

	IdLiveGame uint `gorm:"column:idlivegame"`
	IdTeam uint `gorm:"column:idteam"`

	IdRosterGoal uint `gorm:"column:idroster_goal"`
	IdRosterA1 uint `gorm:"column:idroster_a1"`
	IdRosterA2 uint `gorm:"column:idroster_a2"`

	GoalType string `gorm:"column:goaltype"`
	Period   uint
	TimeScored string `gorm:"column:timescored"`

	// Controll information
	ID uint `gorm:"column:idscoring;primary_key"`
	IsActive bool    `gorm:"column:isactive"`
	UpdatedAt  time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt  time.Time

}