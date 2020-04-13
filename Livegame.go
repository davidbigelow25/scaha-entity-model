package entity

import (
	"time"
)

//
// This represents a Club
//
// ToDo:  Adding all the staff Relationships
//
// ToDo: Add all the teams Here

type Livegame struct {
	IdTeamHome uint  `gorm:"column:idteamhome"`
	IdTeamAway uint  `gorm:"column:idteamaway"`
	HomeTeam   Teams `gorm:"foreignkey:IdTeamHome"`
	AwayTeam   Teams `gorm:"foreignkey:IdTeamAway"`

	HomeShotsOnGoal []*Sog     `gorm:"foreignkey:IdTeam;association_foreignkey:IdTeamHome"`
	AwayShotsOnGoal []*Sog     `gorm:"foreignkey:IdTeam;association_foreignkey:IdTeamAway"`
	HomeScoring     []*Scoring `gorm:"foreignkey:IdTeam;association_foreignkey:IdTeamHome"`
	AwayScoring     []*Scoring `gorm:"foreignkey:IdTeam;association_foreignkey:IdTeamAway"`
	HomePenalties   []*Penalty `gorm:"foreignkey:IdTeam;association_foreignkey:IdTeamHome"`
	AwayPenalties   []*Penalty `gorm:"foreignkey:IdTeam;association_foreignkey:IdTeamAway"`
	Mia             []*Mia     `gorm:"foreignkey:IdLiveGame"`

	IdSchedule uint      `gorm:"column:idschedule"`
	Schedule   Schedule  `gorm:"foreignkey:IdSchedule"`
	TypeTag    string    `gorm:"column:typetag"`
	StateTag   string    `gorm:"column:statetag"`
	VenueTag   string    `gorm:"column:venuetag"`
	Venue      Venue     `gorm:"foreignkey:Tag;association_foreignkey:VenueTag"`
	SheetName  string    `gorm:"column:sheetname"`
	ActDate    time.Time `gorm:"column:actdate"`
	StartTime  string    `gorm:"column:starttime"`
	GameTag    string    `gorm:"column:gametag"`
	ScoreHome  uint      `gorm:"column:scorehome"`
	ScoreAway  uint      `gorm:"column:scoreaway"`
	GameNotes  string    `gorm:"column:gamenotes"`

	ID        uint      `gorm:"column:idlivegame;primary_key"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt     *time.Time `sql:"index"`
	CreatedAt time.Time
}
