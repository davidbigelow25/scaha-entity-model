package entity

import "github.com/jinzhu/gorm"

type NewsItem struct {
	gorm.Model
	Title    string
	Author   string
	Subject  string
	Intro    string
	Body     string
	State    string
	IsActive bool
}
