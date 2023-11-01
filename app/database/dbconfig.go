package DBConn

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
)

var hErr = func(ctx string, err error) {
	if err != nil {
		log.Fatal(fmt.Errorf(ctx, err))
	}
}

func InitializeDB(user, pass, addr, dbName string, locker *sync.Once) (_db *sql.DB) {
	cfg := mysql.Config{
		User:   user,
		Passwd: pass,
		Net:    "tcp",
		Addr:   addr, // host:port
		DBName: dbName,
	}

	locker.Do(func() {
		// TODO: figure out which timeout is better to wait
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var err error

		_db, err = sql.Open("mysql", cfg.FormatDSN())
		defer func() {
			hErr("on database setup", err)
		}()

		hErr("to ping server", _db.PingContext(ctx))
	})
	return
}
