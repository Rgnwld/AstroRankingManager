package AstroRoutes

import (
	"Astro/webscrapper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyNewsRoutes(router *gin.RouterGroup) {
	router.GET("/news", GetNews)
}

func hErr(er error) {
	fmt.Println(er)
}

func GetNews(c *gin.Context) {

	news, err := webscrapper.FetchNews()

	if err != nil {
		hErr(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, news)
}
