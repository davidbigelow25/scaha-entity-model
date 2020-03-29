package entity

import "time"

type Teams struct {

	Idclub  uint
	Sname string
	Teamname string
	Teamgender string
	Idskilllevel uint
	Iddivisions uint
	Divisiontag string
	Isexhibition bool
	Year 	  uint
	Seasontag string
	Scheduletags  string
	Grouptag string
	PlayerCountCalc uint `gorm:"column:playercountcalc"`

	// Controll information
	ID uint `gorm:"column:idteams;primary_key"`
	IsActive bool    `gorm:"column:isactive"`
	UpdatedAt  time.Time `gorm:"column:updated"`
//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt  time.Time
}