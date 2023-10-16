package token

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Create the JWT key used to create the signature
var JWTKey = []byte("my_secret_key")

// For simplification, we're storing the users information as an in-memory map in our code
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Create the GetToken handler
func GetToken(creds Credentials) string {

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		fmt.Println("Error")
		return ""
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
		return ""
	}

	return tokenString
}

func AuthenticatedAction() func(c *gin.Context) {
	return func(c *gin.Context) {
		tknStr := c.Query("token")

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
			return JWTKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{
					"response": "Err Signature Invalid",
				})
				c.Abort()
			}
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"response": "Bad Request",
			})
			c.Abort()
		}

		if !tkn.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"response": "Not Authorized",
			})
			c.Abort()
		}
	}

}
