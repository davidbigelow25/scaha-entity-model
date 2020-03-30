package entity

import "time"

type Sog struct {

	IdLiveGame uint `gorm:"column:idlivegame;primary_key"`
	IdTeam uint `gorm:"column:idteam:primary_key"`
	IdRoster uint `gorm:"column:idroster:primary_key"`

	// What is the playtime in minutes and seconds
	PlayTime string `gorm:"column:playtime"`

	// Shots for 9 periods if it comes to that
	Shots1  uint
	Shots2  uint
	Shots3  uint
	Shots4  uint
	Shots5  uint
	Shots6  uint
	Shots7  uint
	Shots8  uint
	Shots9  uint

	// Controll information
	ID uint `gorm:"column:idsog;primary_key"`
	IsActive bool    `gorm:"column:isactive"`
	UpdatedAt  time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt  time.Time

}