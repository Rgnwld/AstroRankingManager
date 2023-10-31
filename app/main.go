package main

import (
	DBConn "Astro/database"
	AstroRoutes "Astro/routes"
	Token "Astro/token"
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	dbRankingUser = os.Getenv("DB_USER")
	dbRankingPass = os.Getenv("DB_PASS")
	dbRankingAddr = os.Getenv("DB_ADDR")
	dbRankingName = os.Getenv("DB_NAME")

	dbMigrationUser = os.Getenv("DB_USER")
	dbMigrationPass = os.Getenv("DB_PASS")
	dbMigrationAddr = os.Getenv("DB_ADDR")
	dbMigrationName = os.Getenv("DB_NAME")
)

var (
	migrationDbOnce sync.Once
	rankingDbOnce   sync.Once
)

func initRankingDb() *sql.DB {
	var db *sql.DB
	rankingDbOnce.Do(func() {
		db = DBConn.InitializeDB(
			dbRankingUser,
			dbRankingPass,
			dbRankingAddr,
			dbRankingName,
		)
	})
	return db
}

func init() {
	migrationDbOnce.Do(func() {
		db := DBConn.InitializeDB(
			dbMigrationUser,
			dbMigrationPass,
			dbMigrationAddr,
			dbMigrationName,
		)
		DBConn.MigrationUp(db)
	})
}

func main() {
	router := gin.Default()

	rankingDb := initRankingDb()
	rankingRepo := DBConn.NewRankingRepository(rankingDb)

	authRoutes := router.Group("/auth") //Routes for Authentication

	authenticatedRoutes := router.Group("/v1") // Authenticated Route
	authenticatedRoutes.Use(Token.AuthenticatedAction())

	//region Routes
	//Public
	AstroRoutes.TokenRoutes(authRoutes)
	AstroRoutes.UserRoutes(authRoutes)

	//Authenticated
	AstroRoutes.RankingRoutes(authenticatedRoutes, rankingRepo)
	//endregion

	log.Fatal(router.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
