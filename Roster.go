package entity

import "time"

type Roster struct {
	IdPerson uint `gorm:"column:idperson"`
	IdTeam   uint `gorm:"column:idteam"`

	JerseyNumber string `gorm:"column:jerseynumber"`
	Year         uint
	SeasonTag    string `gorm:"column:seasontag"`
	Notes        string
	RosterType   string `gorm:"column:rostertype"`
	//
	// Some key Dates here
	Added      time.Time
	Eligible   time.Time
	Dropped    time.Time
	RosterDate time.Time `gorm:"column:rosterdate"`

	IsPlayerUp  BitBool `gorm:"column:isplayerup"`
	IsVoid      BitBool `gorm:"column:isvoid"`
	IsGoalie    uint    `gorm:"column:isgoalie"`
	IsConfirmed BitBool `gorm:"column:isconfirmed"`
	IsSupsended BitBool `gorm:"column:issuspended"`

	Person Person `gorm:"foreignkey:IdPerson"`
	Player Player `gorm:"foreignkey:IdPerson;association_foreignkey:IdPerson"`

	// Controll information
	ID        uint      `gorm:"column:idroster;primary_key"`
	IsActive  *bool     `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}
