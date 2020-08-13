package entity

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Person struct {
	FirstName        string `gorm:"size:45;column:fname;not null"`
	LastName         string `gorm:"size:45;column:lname;not null"`
	Email            string `gorm:"size:45;not null"`
	Phone            string `gorm:"size:14"`
	Address          string `gorm:"size:75"`
	City             string `gorm:"size:45"`
	State            string `gorm:"size:2"`
	Zipcode          int
	Country          string `gorm:"size:45"`
	Gender           string `gorm:"size:1"`
	Dob              time.Time
	Citizenship      string   `gorm:"size:45"`
	BirthCertificate BitBool  `gorm:"column:birthcertificate"`
	IdProfile        uint     `gorm:"column:idprofile;not null"`
	Profile          *Profile `gorm:"foreignkey:idprofile;PRELOAD:false"`
	IdFamily         int
	Family           *Family     `gorm:"foreignkey:IdPerson"`
	UsaHockeys       []UsaHockey `gorm:"foreignkey:IdPerson"`

	// Controll information
	ID        uint      `gorm:"column:idperson;primary_key"`
	IsActive  *bool     `gorm:"column:isactive"`
	UpdatedAt time.Time `gorm:"column:updated"`
	//	DeletedAt  *time.Time `sql:"index"`
	CreatedAt time.Time
}

//
// Hey, lets be smart and filter out all the garbaded that can come it
// make it html safeish
func (p *Person) Prepare() {
	p.ID = 0
	p.FirstName = html.EscapeString(strings.TrimSpace(p.FirstName))
	p.LastName = html.EscapeString(strings.TrimSpace(p.LastName))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

//
// Here are all the validation rules we can apply here for the given object
//
func (p *Person) Validate() error {

	if p.FirstName == "" {
		return errors.New("first name required")
	}
	if p.LastName == "" {
		return errors.New("last name")
	}
	return nil
}
