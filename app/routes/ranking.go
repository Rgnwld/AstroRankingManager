package AstroRoutes

import (
	DBConn "Astro/repository"
	Token "Astro/token"
	AstroTypes "Astro/types"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RankingHandlers struct {
	rankingRepo DBConn.RankingRepository
}

func ApplyRankingRoutes(router *gin.RouterGroup, rankingRepo DBConn.RankingRepository) {
	hs := RankingHandlers{rankingRepo: rankingRepo}

	router.GET("/ranking/:mapId", hs.mapAllRankingsByMap)
	router.GET("/ranking", hs.mapPlayerAllRankings)
	router.POST("/ranking", hs.newRanking)
	router.PATCH("/ranking/:id", hs.updateRanking)
	router.DELETE("/ranking/:id", hs.removeRanking)
}

// region GET

// TestRoute

func (ah *RankingHandlers) mapPlayerAllRankings(c *gin.Context) {
	//List selected player ranks by time
	//Order it by map
	tknStr := c.Query("token")

	_, claims, err := Token.ParseToken(tknStr)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Token",
		})
	}

	times, err := ah.rankingRepo.GetPlayerAllRanking(c.Request.Context(), claims.UserId)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, times)
}

func (ah *RankingHandlers) mapAllRankingsByMap(c *gin.Context) {
	mapId := c.Param("mapId")

	times, err := ah.rankingRepo.GetRankingByMap(c.Request.Context(), mapId)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, times)
}

// endregion

// region POST
func (ah *RankingHandlers) newRanking(c *gin.Context) {
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

	err = ah.rankingRepo.AddRanking(c.Request.Context(), newRankMetadata)
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
func (ah *RankingHandlers) updateRanking(c *gin.Context) {

	id := c.Param("id")
	var updateRank AstroTypes.TimeObj

	if err := c.BindJSON(&updateRank); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong. \nPlease, check your request."})
		return
	}

	result, err := ah.rankingRepo.PatchRankingTime(c.Request.Context(), id, updateRank)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ID not founded"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Time Updated from: " + strconv.Itoa(result.TimeInSeconds) + " to: " + strconv.Itoa(updateRank.TimeInSeconds)})
}

//endregion

// region DELETE
func (ah *RankingHandlers) removeRanking(c *gin.Context) {

	id := c.Param("id")

	result, err := ah.rankingRepo.DeleteRanking(c.Request.Context(), id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ID not founded"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result.Id + " was removed"})
}

//endregion
