package AstroRoutes

import (
	DBConn "Astro/database"
	AstroTypes "Astro/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RankingRoutes(router *gin.RouterGroup) {

	router.GET("/ranking/:mapId", mapAllRankingsByMap)
	router.GET("/ranking", mapAllRankingsByPlayer)
	router.POST("/ranking", newRanking)
	router.PATCH("/ranking/:id", updateRanking)
	router.DELETE("/ranking/:id", removeRanking)
}

// region GET

//TestRoute
func mapAllRankingsByMap(c *gin.Context) {
	times := DBConn.GetRankings()

	c.IndentedJSON(http.StatusOK, times)
}

func mapAllRankingsByMap(c *gin.Context) {
	idQuery := c.Param("mapId")
	times := DBConn.GetRankings()

	c.IndentedJSON(http.StatusOK, times)
}

func mapAllRankingsByPlayer(c *gin.Context) {
	//List selected player ranks by time
	//Order it by map
	idQuery := c.Param("id")
	times := DBConn.GetSpecificRanking(idQuery)

	c.IndentedJSON(http.StatusOK, times)
}

// endregion

// region POST
func newRanking(c *gin.Context) {
	var newRank AstroTypes.TimeObj

	tkn, claims, err := Token.ParseToken(c.Query("token"))

	if err != nil{
		.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
	}

	if err := c.BindJSON(&newRank); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Unexpected Body",
		})
		return
	}

	newRankMetadata := AstroType.UserTimeObj{
		"timeInSeconds": newRank.timeInSeconds, 
		"map": newRank.map, 
		username 
	}

	DBConn.AddRanking(newRank)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "ranking was setted",
		"object":  newRank,
	})
}

//endregion

// region PATCH
func updateRanking(c *gin.Context) {

	id := c.Param("id")
	var updateRank AstroTypes.TimeObj

	if err := c.BindJSON(&updateRank); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong. \nPlease, check your request."})
		return
	}

	result, err := DBConn.PatchRankingTime(id, updateRank)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ID not founded"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Time Updated from: " + strconv.Itoa(result.TimeInSeconds) + " to: " + strconv.Itoa(updateRank.TimeInSeconds)})
}

//endregion

// region DELETE
func removeRanking(c *gin.Context) {

	id := c.Param("id")

	result, err := DBConn.DeleteRanking(id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ID not founded"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result.Id + " was removed"})
}

//endregion
