package AstroTypes

import "github.com/golang-jwt/jwt/v5"

// Create a struct to read the username and password from the request body
type DBCredentials struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
}

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
