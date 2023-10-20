package AstroRoutes

import (
	DBConn "Astro/database"
	Token "Astro/token"
	AstroTypes "Astro/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/login", logIn)
	router.POST("/signin", signIn)
}

func logIn(c *gin.Context) {
	var cred AstroTypes.Credentials

	if err := c.BindJSON(&cred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	dbuser, err := DBConn.GetUser(cred.Username)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "User: " + cred.Username + " not Found",
		})
		return
	}

	if !CheckPasswordHash(cred.Password, dbuser.HashedPassword) || err != nil {

		fmt.Sprintln(err)

		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	token := Token.GetToken(dbuser)

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func signIn(c *gin.Context) {
	var cred AstroTypes.Credentials

	if err := c.BindJSON(&cred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"response": http.StatusBadRequest,
			"body":     c.Request.Body,
		})

		c.Abort()
		return
	}

	hashedPassword, err := HashPassword(cred.Password)
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"response": http.StatusInternalServerError,
			"body":     c.Request.Body,
		})

		return
	}

	userid := uuid.NewString()

	hashedCred := AstroTypes.DBCredentials{
		Id:             userid,
		Username:       cred.Username,
		HashedPassword: hashedPassword,
	}

	DBConn.CreateUser(hashedCred)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User was created",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
