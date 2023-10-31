package DBConn

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var hErr = func(ctx string, err error) {
	if err != nil {
		log.Fatal(fmt.Errorf(ctx, err))
	}
}

func InitializeDB(user, pass, addr, dbName string) *sql.DB {
	cfg := mysql.Config{
		User:   user,
		Passwd: pass,
		Net:    "tcp",
		Addr:   addr, // host:port
		DBName: dbName,
	}

	// "root:astropass@tcp(127.0.0.1:3306)/"
	_db, err := sql.Open("mysql", cfg.FormatDSN())
	defer func() {
		hErr("on database setup", err)

		_ = _db.Close()
	}()
	return _db
}

func MigrationUp(_db *sql.DB) {

	tx, err := _db.Begin()
	hErr("on init transaction", err)

	defer hErr("on init transaction", tx.Commit())

	_, err = tx.Exec(`
		CREATE DATABASE IF NOT EXISTS AstroRankings;

	   	USE AstroRankings;
	   	    
		CREATE TABLE IF NOT EXISTS userRanking ( id varchar(64), userId varchar(64) , timeInSeconds integer, mapId integer);

		CREATE TABLE IF NOT EXISTS userTable ( id varchar(64), username varchar(32), hashedpassword varchar(64));
	`)
}
