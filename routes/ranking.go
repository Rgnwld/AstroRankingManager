package AstroRoutes

import (
	DBConn "Astro/database"
	AstroTypes "Astro/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RankingRoutes(router *gin.Engine) {

	router.GET("/ranking", mapAllRankings)
	router.GET("/ranking/:id", mapOneRanking)
	router.POST("/ranking", newRanking)
	router.PATCH("/ranking/:id", updateRanking)
	router.DELETE("/ranking/:id", removeRanking)
}

// region GET
func mapAllRankings(c *gin.Context) {

	times := DBConn.GetRankings()

	c.IndentedJSON(http.StatusOK, times)
}

func mapOneRanking(c *gin.Context) {

	idQuery := c.Param("id")
	times := DBConn.GetSpecificRanking(idQuery)

	c.IndentedJSON(http.StatusOK, times)
}

// endregion

// region POST
func newRanking(c *gin.Context) {
	var newRank AstroTypes.TimeObj

	if err := c.BindJSON(&newRank); err != nil {
		return
	}

	DBConn.AddRanking(newRank)

	c.IndentedJSON(http.StatusCreated, newRank)
}

//endregion

// region PATCH
func updateRanking(c *gin.Context) {

	id := c.Param("id")
	var updateRank AstroTypes.TimeObj

	if err := c.BindJSON(&updateRank); err != nil {
		return
	}

	DBConn.PatchRankingTime(id, updateRank)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Time Updated to: " + strconv.Itoa(updateRank.TimeInSeconds)})
}

//endregion

// region DELETE
func removeRanking(c *gin.Context) {

	id := c.Param("id")

	DBConn.DeleteRanking(id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": id + " was removed"})
}

//endregion
