package token

import (
	"fmt"
	"net/http"
	"time"

	AstroTypes "Astro/types"

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

// Create the GetToken handler
func GetToken(creds AstroTypes.Credentials) string {

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		fmt.Println("Error")
		return ""
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &AstroTypes.Claims{
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
		claims := &AstroTypes.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
			return JWTKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{
					"response": "Err Signature Invalid",
				})
				c.Abort()
				return
			}
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"response": "Bad Request",
			})
			c.Abort()
			return
		}

		if !tkn.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"response": "Not Authorized",
			})
			c.Abort()
			return
		}
	}

}

func ParseToken(tknStr string) (*jwt.Token, *AstroTypes.Claims,  error) {
	// Initialize a new instance of `Claims`
	claims := &AstroTypes.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return JWTKey, nil
	})

	return tkn, claims, err
}