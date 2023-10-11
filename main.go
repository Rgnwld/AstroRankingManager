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
	AstroRoutes.GetRanking(router)         //Get "/ranking"
	AstroRoutes.GetSpecificRanking(router) //Get "/ranking/:id"
	AstroRoutes.AddRanking(router)         //Post "/ranking" + TimeObj
	AstroRoutes.PatchRanking(router)       //Patch "/ranking/:id" + TimeObj
	AstroRoutes.DeleteRanking(router)      //Patch "/ranking/:id"
	//endregion

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
