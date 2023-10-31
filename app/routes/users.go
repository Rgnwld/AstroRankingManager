package AstroRoutes

import (
	DBConn "Astro/repository"
	Token "Astro/token"
	AstroTypes "Astro/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type usersHandlers struct {
	userRepo DBConn.UserRepository
}

func ApplyUserRouters(userRepo DBConn.UserRepository, router *gin.RouterGroup) {
	uh := &usersHandlers{userRepo: userRepo}

	router.POST("/login", uh.logIn)
	router.POST("/signin", uh.signIn)
}

func (uh *usersHandlers) logIn(c *gin.Context) {
	var cred AstroTypes.Credentials

	if err := c.BindJSON(&cred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "fail to deserialize",
		})
		return
	}

	dbuser, err := uh.userRepo.GetUser(c.Request.Context(), cred.Username)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("User: [%s] not Found", cred.Username),
		})
		return
	}

	if !CheckPasswordHash(cred.Password, dbuser.HashedPassword) || err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": Token.GetToken(dbuser),
	})
}

func (uh *usersHandlers) signIn(c *gin.Context) {
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

	err = uh.userRepo.CreateUser(c.Request.Context(), hashedCred)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User was created",
	})
}

// TODO: move to cipher package in order to decouple and testing stuff
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
