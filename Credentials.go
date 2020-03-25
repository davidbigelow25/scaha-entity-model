package entity


// Create a struct that models the structure of a user, both in the request body, and in the DB
// This is the basic structure that will be used
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

