package entity

import (
	"github.com/jinzhu/gorm"
)

//
// This represents a Club
//
// ToDo:  Adding all the staff Relationships
//
// ToDo: Add all the teams Here

type Club struct {
	gorm.Model
	CahaNumber    string
	Tag  		  string
	ShortName 	  string
	Name 		  string
	Url           string
	IsActive 	  *bool
	IsHighSchool  BitBool
	MobileImageURL string
}

