package AstroRoutes

import (
	DBConn "Astro/database"
	AstroTypes "Astro/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// region GET
func mapAllRankings(c *gin.Context) {

	times := DBConn.GetRankings()

	c.IndentedJSON(http.StatusOK, times)
}

func GetRanking(router *gin.Engine) {

	router.GET("/ranking", mapAllRankings)
}

func mapOneRanking(c *gin.Context) {

	idQuery := c.Param("id")
	times := DBConn.GetSpecificRanking(idQuery)

	c.IndentedJSON(http.StatusOK, times)
}

func GetSpecificRanking(router *gin.Engine) {
	router.GET("/ranking/:id", mapOneRanking)
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

func AddRanking(router *gin.Engine) {

	router.POST("/ranking", newRanking)
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

func PatchRanking(router *gin.Engine) {

	router.PATCH("/ranking/:id", updateRanking)
}

//endregion

// region DELETE
func removeRanking(c *gin.Context) {

	id := c.Param("id")

	DBConn.DeleteRanking(id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": id + " was removed"})
}

func DeleteRanking(router *gin.Engine) {

	router.DELETE("/ranking/:id", removeRanking)
}

//endregion
