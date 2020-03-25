package scaha_orm

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	RoleName        string
	RoleDescription string
	IdParentRole    int
	DefaultRole     int
	InheritedRoles  Roles `gorm:"foreignkey:IdParentRole;"`
}
