package cmd

import (
	DBConn "Astro/database"
	"Astro/repository"
	AstroRoutes "Astro/routes"
	Token "Astro/token"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	rankingDbOnce sync.Once
	usersDbOnce   sync.Once
)

func (app *app) serveApi() error {
	router := gin.Default()

	router.Use(CORSMiddleware())

	rankingDb := DBConn.InitializeDB(
		app.config.rankingDbOpts.user,
		app.config.rankingDbOpts.pass,
		app.config.rankingDbOpts.addr,
		app.config.rankingDbOpts.dbName,
		&rankingDbOnce)

	usersDb := DBConn.InitializeDB(
		app.config.usersDbOpts.user,
		app.config.usersDbOpts.pass,
		app.config.usersDbOpts.addr,
		app.config.usersDbOpts.dbName,
		&usersDbOnce)

	rankingRepo := repository.NewRankingRepository(rankingDb)
	usersRepo := repository.NewUserRepository(usersDb)

	publicRoutes := router.Group("/public")
	authRoutes := router.Group("/auth")        // Routes for Authentication
	authenticatedRoutes := router.Group("/v1") // Authenticated Route

	authenticatedRoutes.Use(Token.AuthenticatedAction())

	//region Routes
	//Public
	AstroRoutes.TokenRoutes(authRoutes)
	AstroRoutes.ApplyUserRouters(usersRepo, authRoutes)
	AstroRoutes.ApplyNewsRoutes(publicRoutes)

	//Authenticated
	AstroRoutes.ApplyRankingRoutes(authenticatedRoutes, rankingRepo)
	//endregion

	return router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
