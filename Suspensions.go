package entity

import "time"

type Suspensions struct {
	IdPerson uint `gorm:"column:idperson"`
	IdClub   uint `gorm:"column:idclub"`
	IdTeam   uint `gorm:"column:idteam"`

	SuspensionDate time.Time `gorm:"column:suspensiondate"`

	Infraction    string
	NumberOfGames uint    `gorm:"column:numberofgames"`
	IsMatch       BitBool `gorm:"column:ismatch"`

	Eligibility string
	ReviewFile  string

	IsServed BitBool `gorm:"column:isserved"`

	// Controll information
	ID        uint      `gorm:"column:idsuspensions;primary_key;AUTO_INCREMENT"`
	IsActive  *bool      `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}
