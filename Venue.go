package entity

import "time"

type Venue struct {
	ShortName   string `gorm:"column:shortname"`
	Description string `gorm:"column:description"`
	GmapParms   string  `gorm:"column:gmapparms"`

	// Control information
	ID        uint `gorm:"column:idvenue;primary_key;AUTO_INCREMENT"`
	Tag       string
	IsActive  *bool      `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}
