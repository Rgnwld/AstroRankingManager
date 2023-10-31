package DBConn

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
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

	// TODO: figure out which timeout is better to wait
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_db, err := sql.Open("mysql", cfg.FormatDSN())
	defer func() {
		hErr("on database setup", err)

		_ = _db.Close()
	}()

	hErr("to ping server", _db.PingContext(ctx))
	return _db
}

func MigrationUp(_db *sql.DB) {
	// TODO: figure out which timeout is better to wait
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	tx, err := _db.Begin()
	hErr("on init transaction", err)

	defer hErr("on commit transaction", tx.Commit())

	_, err = tx.ExecContext(ctx, `
		CREATE DATABASE IF NOT EXISTS AstroRankings;

	   	USE AstroRankings;
	   	    
		CREATE TABLE IF NOT EXISTS userRanking ( id varchar(64), userId varchar(64) , timeInSeconds integer, mapId integer);

		CREATE TABLE IF NOT EXISTS userTable ( id varchar(64), username varchar(32), hashedpassword varchar(64));
	`)
}
