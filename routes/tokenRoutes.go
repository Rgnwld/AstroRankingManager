package AstroRoutes

import (
	Token "Astro/token"
	AstroTypes "Astro/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TokenRoutes(router *gin.RouterGroup) {

	router.GET("/token", TestToken)
}

func TestToken(c *gin.Context) {

	tknStr := c.Query("token")

	// Initialize a new instance of `Claims`
	claims := &AstroTypes.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return Token.JWTKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "Err Signature Invalid",
			})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if !tkn.Valid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Welcome " + claims.Username,
	})
}
