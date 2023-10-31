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

type ApiHandlers struct {
	rankingRepo *DBConn.RankingRepository
}

func RankingRoutes(router *gin.RouterGroup, rankingRepo *DBConn.RankingRepository) {
	hs := ApiHandlers{rankingRepo: rankingRepo}

	router.GET("/ranking/:mapId", hs.mapAllRankingsByMap)
	router.GET("/ranking", hs.mapPlayerAllRankings)
	router.POST("/ranking", hs.newRanking)
	router.PATCH("/ranking/:id", hs.updateRanking)
	router.DELETE("/ranking/:id", hs.removeRanking)
}

// region GET

// TestRoute

func (ah *ApiHandlers) mapPlayerAllRankings(c *gin.Context) {
	//List selected player ranks by time
	//Order it by map
	tknStr := c.Query("token")

	_, claims, err := Token.ParseToken(tknStr)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Token",
		})
	}

	times, err := ah.rankingRepo.GetPlayerAllRanking(claims.UserId)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, times)
}

func (ah *ApiHandlers) mapAllRankingsByMap(c *gin.Context) {
	mapId := c.Param("mapId")

	times, err := ah.rankingRepo.GetRankingByMap(mapId)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, times)
}

// endregion

// region POST
func (ah *ApiHandlers) newRanking(c *gin.Context) {
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

	err = ah.rankingRepo.AddRanking(newRankMetadata)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "ranking was setted",
		"object":  newRank,
	})
}

//endregion

// region PATCH
func (ah *ApiHandlers) updateRanking(c *gin.Context) {

	id := c.Param("id")
	var updateRank AstroTypes.TimeObj

	if err := c.BindJSON(&updateRank); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong. \nPlease, check your request."})
		return
	}

	result, err := ah.rankingRepo.PatchRankingTime(id, updateRank)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ID not founded"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Time Updated from: " + strconv.Itoa(result.TimeInSeconds) + " to: " + strconv.Itoa(updateRank.TimeInSeconds)})
}

//endregion

// region DELETE
func (ah *ApiHandlers) removeRanking(c *gin.Context) {

	id := c.Param("id")

	result, err := ah.rankingRepo.DeleteRanking(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ID not founded"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result.Id + " was removed"})
}

//endregion
