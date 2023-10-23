package DBConn

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var Cfg = mysql.Config{
	User:   "root",
	Passwd: "astropass",
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "AstroRankings",
}

func InitializeDB() {

	_db, err := sql.Open("mysql", "root:astropass@tcp(127.0.0.1:3306)/")
	defer _db.Close()

	_, err = _db.Exec("CREATE database AstroRankings")
	if err != nil {
		fmt.Printf("Database already created\n\n")
		return
	}

	_, err = _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	_, err = _db.Exec("CREATE TABLE userRanking ( id varchar(64), userId varchar(64) , timeInSeconds integer, mapId integer);")
	if err != nil {
		panic(err)
	}

	_, err = _db.Exec("CREATE TABLE userTable ( id varchar(64), username varchar(32), hashedpassword varchar(64));")
	if err != nil {
		panic(err)
	}
}
