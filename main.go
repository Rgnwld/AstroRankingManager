package main

import (
	DBConn "Astro/database"
	AstroRoutes "Astro/routes"
	Token "Astro/token"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	DBConn.InitializeDB()

	authRoutes := router.Group("/auth") //Routes for Authentication

	authenticatedRoutes := router.Group("/v1") // Authenticated Route
	authenticatedRoutes.Use(Token.AuthenticatedAction())

	//region Routes
	//Public
	AstroRoutes.TokenRoutes(authRoutes)

	//Authenticated
	AstroRoutes.RankingRoutes(authenticatedRoutes)
	//endregion

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
