package entity

import "gopkg.in/dgrijalva/jwt-go.v3"

// Create a struct that models the structure of a user, both in the request body, and in the DB
// This is the basic structure that will be used
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// This is custom claims to to pass around all the role information tied to the user when the user authenticates
//
type Claims struct {
	UserName string `json:"username"`
	UserId  uint    `json:"userid"`
	ProfileId uint   `json:"profileid"`
	Roles map[string]bool  `json:"rolemap"`
	jwt.StandardClaims
}

