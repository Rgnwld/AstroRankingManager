package main

import (
	DBConn "Astro/database"
	AstroRoutes "Astro/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	DBConn.InitializeDB()

	//region Routes
	AstroRoutes.RankingRoutes(router)
	AstroRoutes.TokenRoutes(router)
	//endregion

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
