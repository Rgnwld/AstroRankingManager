package AstroRoutes

import (
	Token "Astro/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TokenRoutes(router *gin.RouterGroup) {

	router.GET("/token", TestToken)
}

func TestToken(c *gin.Context) {

	tknStr := c.Query("token")

	tkn, claims, err := Token.ParseToken(tknStr)

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
