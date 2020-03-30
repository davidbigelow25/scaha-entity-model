package entity

import "time"

type Schedule struct {

	Description string
	ScheduleWeekTag string  `gorm:"column:scheduleweektag"`
	SeasonTag string  `gorm:"column:seasontag"`
	GameTag string  `gorm:"column:gametag"`
	DivisionName string `gorm:"column:divsname"`
	TeamCount uint `gorm:"column:teamcount"`
	StartDate time.Time `gorm:"column:startdate"`
	EndDate time.Time `gorm:"column:enddate"`
	ScheduleType string  `gorm:"column:scheduletype"`

	//
	// We can add all the extra needed columns later

	// Controll information
	ID uint `gorm:"column:idschedule;primary_key"`
	Tag string `gorm:"index"`
	IsActive bool    `gorm:"column:isactive"`
	UpdatedAt  time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt  time.Time
}
