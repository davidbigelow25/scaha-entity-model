package entity

import (
	"database/sql/driver"
	"errors"
)

// BitBool is an implementation of a bool for the MySQL type BIT(1).
// This type allows you to avoid wasting an entire byte for MySQL's boolean type TINYINT.
type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (b *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}



// BitBool is an implementation of a bool for the MySQL type BIT(1).
// This type allows you to avoid wasting an entire byte for MySQL's boolean type TINYINT.
type IntBool bool

// Value implements the driver.Valuer interface,
// and turns the IntBool into a bitfield (BIT(1)) for MySQL storage.
func (b IntBool) Value() (driver.Value, error) {
	if b {
		return 1, nil
	} else {
		return 0, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a IntBool
func (b *IntBool) Scan(src interface{}) error {
	v, ok := src.(uint)
	if !ok {
		return errors.New("bad Value Type")
	}
	*b = v == 1
	return nil
}



type Roles []Role

//
// We have to flatten out the roles so the user gets ALL direct and interited roles
// we do not care how the app uses them
//
func (r Roles) Flatten() *Roles {

	var flatten Roles

	for _, role := range r {
		fr := Role{role.Model, role.RoleName, role.RoleDescription, role.IdParentRole, role.DefaultRole, nil}
		flatten = append(flatten, fr)
		if role.InheritedRoles != nil {
			flatten = append(flatten, *role.InheritedRoles.Flatten()...)
		}
	}
	return &flatten
}

//
// We have to flatten out the roles so the user gets ALL direct and interited roles
// we do not care how the app uses them
//
func (r Roles) FlattenAndMap() map[string]bool {

	var flatten Roles

	for _, role := range r {
		fr := Role{role.Model, role.RoleName, role.RoleDescription, role.IdParentRole, role.DefaultRole, nil}
		flatten = append(flatten, fr)
		if role.InheritedRoles != nil {
			flatten = append(flatten, *role.InheritedRoles.Flatten()...)
		}
	}

	var answer = map[string]bool{}
	for _, role := range flatten {
		answer[role.RoleName] = true
	}

	return answer
}
