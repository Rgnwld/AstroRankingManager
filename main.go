package main

import (
	AstroRoutes "Astro/routes"
	astrotypes "Astro/types"

	"github.com/gin-gonic/gin"
)

var testobj = astrotypes.TimeObj{
	Id:            "avbc",
	Username:      "Regis",
	TimeInSeconds: 1,
	Map:           1,
}

func main() {
	router := gin.Default()

	//region Routes
	AstroRoutes.GetRanking(router)         //Get "/ranking"
	AstroRoutes.GetSpecificRanking(router) //Get "/ranking/:id"
	AstroRoutes.AddRanking(router)         //Post "/ranking" + TimeObj
	AstroRoutes.PatchRanking(router)       //Patch "/ranking/:id" + TimeObj
	AstroRoutes.DeleteRanking(router)      //Patch "/ranking/:id"
	//endregion

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
