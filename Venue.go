package entity

import "time"

type Venue struct {
	ShortName   string `gorm:"column:shortname"`
	Description string `gorm:"column:description"`

	// Control information
	ID        uint `gorm:"column:idvenue;primary_key"`
	Tag       string
	IsActive  bool      `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}
