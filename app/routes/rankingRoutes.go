package AstroRoutes

import (
	DBConn "Astro/database"
	Token "Astro/token"
	AstroTypes "Astro/types"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RankingRoutes(router *gin.RouterGroup) {

	router.GET("/ranking/:mapId", mapAllRankingsByMap)
	router.GET("/ranking", mapPlayerAllRankings)
	router.POST("/ranking", newRanking)
	router.PATCH("/ranking/:id", updateRanking)
	router.DELETE("/ranking/:id", removeRanking)
}

// region GET

// TestRoute

func mapPlayerAllRankings(c *gin.Context) {
	//List selected player ranks by time
	//Order it by map
	tknStr := c.Query("token")

	_, claims, err := Token.ParseToken(tknStr)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Token",
		})
	}

	times := DBConn.GetPlayerAllRanking(claims.UserId)
	c.IndentedJSON(http.StatusOK, times)
}

func mapAllRankingsByMap(c *gin.Context) {
	mapId := c.Param("mapId")

	times := DBConn.GetRankingByMap(mapId)
	c.IndentedJSON(http.StatusOK, times)
}

// endregion

// region POST
func newRanking(c *gin.Context) {
	var newRank AstroTypes.TimeObj

	_, claims, err := Token.ParseToken(c.Query("token"))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}

	if err := c.BindJSON(&newRank); err != nil {
		fmt.Println(err)

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Unexpected Body",
		})
		return
	}

	newRankMetadata := AstroTypes.UserTimeObj{
		Id:            uuid.NewString(),
		UserId:        claims.UserId,
		TimeInSeconds: newRank.TimeInSeconds,
		MapId:         newRank.MapId,
	}

	DBConn.AddRanking(newRankMetadata)

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
