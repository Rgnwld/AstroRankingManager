package DBConn

import (
	astrotypes "Astro/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBConnection() *sql.DB {

	db, err := sql.Open("mysql", Cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	return db
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

	_, err = _db.Exec("CREATE TABLE userRanking ( id varchar(32), username varchar(32) , timeInSeconds integer, map integer);")
	if err != nil {
		panic(err)
	}
}

func AddRanking(tobj astrotypes.TimeObj) {

	_db := OpenDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	_, err = _db.Exec("INSERT INTO userRanking ( id, username, timeInSeconds, map)  VALUES (?, ?, ?, ?);",
		tobj.Id, tobj.Username, tobj.TimeInSeconds, tobj.Map)

	if err != nil {
		panic(err)
	}
}

func GetRankings() []astrotypes.TimeObj {

	_db := OpenDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	results, err := _db.Query("SELECT * FROM userRanking")
	if err != nil {
		panic(err)
	}

	var times []astrotypes.TimeObj

	for results.Next() {
		var userTimes astrotypes.TimeObj

		err = results.Scan(&userTimes.Id, &userTimes.Username, &userTimes.TimeInSeconds, &userTimes.Map)

		if err != nil {
			panic(err.Error())
		}

		times = append(times, userTimes)
	}

	return times
}

func GetSpecificRanking(id string) astrotypes.TimeObj {

	_db := OpenDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	results, err := _db.Query("SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		panic(err)
	}

	var times astrotypes.TimeObj

	for results.Next() {
		var userTimes astrotypes.TimeObj

		err = results.Scan(&userTimes.Id, &userTimes.Username, &userTimes.TimeInSeconds, &userTimes.Map)

		if err != nil {
			panic(err.Error())
		}

		times = userTimes
	}

	return times
}

func PatchRankingTime(id string, tobj astrotypes.TimeObj) {

	_db := OpenDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	_, err = _db.Exec("UPDATE userRanking SET timeInSeconds=? WHERE id=?", tobj.TimeInSeconds, id)

	if err != nil {
		panic(err)
	}
}

func DeleteRanking(id string) {

	_db := OpenDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	_, err = _db.Exec("DELETE FROM userRanking WHERE id=?", id)

	if err != nil {
		panic(err)
	}
}
