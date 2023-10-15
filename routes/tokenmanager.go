package AstroRoutes

import (
	Token "Astro/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenRoutes(router *gin.Engine) {

	router.GET("/token")
	router.POST("/token", LogIn)
}

func LogIn(c *gin.Context) {

	var cred Token.Credentials

	if err := c.BindJSON(&cred); err != nil {
		return
	}

	tostr := fmt.Sprintf("%+v", cred)
	println(tostr)

	token := Token.GetToken(cred)

	c.IndentedJSON(http.StatusOK, token)
}

// func AuthenticatedAction(c *gin.Context, action func()) func() {

// 	token := c.Query("token")

// 	if Token.AuthenticateToken(token) {
// 		return action
// 	} else {
// 		print("Token not valid")

// 		return func() {}
// 	}
// }

// func TestToken(c *gin.Context) {

// }
