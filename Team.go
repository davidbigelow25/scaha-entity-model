package entity

import "time"

type Teams struct {
	IdClub uint `gorm:"column:idclub"`

	Name          string `gorm:"column:sname"`
	TeamName      string `gorm:"column:teamname"`
	TeamGender    string `gorm:"column:teamgender"`
	IdSkillLevel  uint   `gorm:"column:IdSkillLevel"`
	IdDivisions   uint   `gorm:"column:iddivisions"`
	DivisionTag   string `gorm:"column:divisiontag"`
	SkillLevelTag string `gorm:"column:skillleveltag"`
	IsExhibition  bool   `gorm:"column:isexhibition"`
	Year          uint
	SeasonTag     string `gorm:"column:seasontag"`
	ScheduleTags  string `gorm:"column:scheduletags"`
	GroupTag      string `gorm:"column:grouptag"`

	//  Rosters and List of Suspensions
	RosterSpots []*Roster      `gorm:"foreignkey:IdTeam"`
	Suspensions []*Suspensions `gorm:"foreignkey:IdTeam"`

	// Controll information
	ID        uint      `gorm:"column:idteams;primary_key"`
	IsActive  *bool      `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}
