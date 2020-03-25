package entity

import "github.com/dgrijalva/jwt-go"

// This is custom claims to to pass around all the role information tied to the user when the user authenticates
//
type Claims struct {
	UserName string `json:"username"`
	UserId  uint    `json:"userid"`
	ProfileId uint   `json:"profileid"`
	Roles map[string]bool  `json:"rolemap"`
	jwt.StandardClaims
}
