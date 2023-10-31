package main

import (
	DBConn "Astro/database"
	"Astro/repository"
	AstroRoutes "Astro/routes"
	Token "Astro/token"
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	dbRankingUser = os.Getenv("DB_RANKING_USER")
	dbRankingPass = os.Getenv("DB_RANKING_PASS")
	dbRankingAddr = os.Getenv("DB_RANKING_ADDR")
	dbRankingName = os.Getenv("DB_RANKING_NAME")

	dbUsersUser = os.Getenv("DB_USERS_USER")
	dbUsersPass = os.Getenv("DB_USERS_PASS")
	dbUsersAddr = os.Getenv("DB_USERS_ADDR")
	dbUsersName = os.Getenv("DB_USERS_NAME")

	dbMigrationUser = os.Getenv("DB_MIGRATION_USER")
	dbMigrationPass = os.Getenv("DB_MIGRATION_PASS")
	dbMigrationAddr = os.Getenv("DB_MIGRATION_ADDR")
	dbMigrationName = os.Getenv("DB_MIGRATION_NAME")
)

var (
	migrationDbOnce sync.Once
	rankingDbOnce   sync.Once
	usersDbOnce     sync.Once
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

func initUsersDb() *sql.DB {
	var db *sql.DB
	usersDbOnce.Do(func() {
		db = DBConn.InitializeDB(
			dbUsersPass,
			dbUsersAddr,
			dbUsersName,
			dbUsersUser,
		)
	})
	return db
}

func init() {
	// apply migrations before run app
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
	rankingRepo := repository.NewRankingRepository(rankingDb)

	usersDb := initUsersDb()
	usersRepo := repository.NewUserRepository(usersDb)

	authRoutes := router.Group("/auth") //Routes for Authentication

	authenticatedRoutes := router.Group("/v1") // Authenticated Route
	authenticatedRoutes.Use(Token.AuthenticatedAction())

	//region Routes
	//Public
	AstroRoutes.TokenRoutes(authRoutes)
	AstroRoutes.ApplyUserRouters(usersRepo, authRoutes)

	//Authenticated
	AstroRoutes.ApplyRankingRoutes(authenticatedRoutes, rankingRepo)
	//endregion

	log.Fatal(router.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
