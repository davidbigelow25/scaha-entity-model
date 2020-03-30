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

	HomeShotsOnGoal Sog `gorm:"foreignkey:IdTeamHome,IdLiveGame;association_foreignkey:IdTeam,IdLiveGame"`
	AwayShotsOnGoal Sog `gorm:"foreignkey:IdTeamAway,IdLiveGame,;association_foreignkey:IdTeam,IdLiveGame"`

//	HomeScoring []*Scoring `gorm:"foreignkey:IdLiveGame,IdTeamHome;association_foreignkey:IdLiveGame,IdTeam"`
//	AwayScoring []*Scoring `gorm:"foreignkey:IdLiveGame,IdTeamAway;association_foreignkey:IdLiveGame,IdTeam"`

	Idschedule uint
	Typetag    string
	Statetag   string
	Venuetag   string
	Sheetname  string
	Actdate    time.Time
	Starttime  string
	Gametag    string
	Scorehome  uint
	Scoreaway  uint
	Gamenotes  string

	ID        uint      `gorm:"column:idlivegame;primary_key"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt     *time.Time `sql:"index"`
	CreatedAt time.Time
}
