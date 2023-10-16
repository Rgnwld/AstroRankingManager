package DBConn

import (
	astrotypes "Astro/types"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func OpenRankingDBConnection() *sql.DB {

	db, err := sql.Open("mysql", Cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	return db
}

func AddRanking(tobj astrotypes.UserTimeObj) {

	_db := OpenUserDBConnection()
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

func GetRankings() []astrotypes.UserTimeObj {

	_db := OpenUserDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	results, err := _db.Query("SELECT * FROM userRanking ORDER BY timeInSeconds ASC")
	if err != nil {
		panic(err)
	}

	var times []astrotypes.UserTimeObj

	for results.Next() {
		var userTimes astrotypes.UserTimeObj

		err = results.Scan(&userTimes.Id, &userTimes.Username, &userTimes.TimeInSeconds, &userTimes.Map)

		if err != nil {
			panic(err.Error())
		}

		times = append(times, userTimes)
	}

	return times
}

func GetSpecificRanking(id string) astrotypes.UserTimeObj {

	_db := OpenUserDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	results, err := _db.Query("SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		panic(err)
	}

	var times astrotypes.UserTimeObj

	for results.Next() {
		var userTimes astrotypes.UserTimeObj

		err = results.Scan(&userTimes.Id, &userTimes.Username, &userTimes.TimeInSeconds, &userTimes.Map)

		if err != nil {
			panic(err.Error())
		}

		times = userTimes
	}

	return times
}

func PatchRankingTime(id string, tobj astrotypes.TimeObj) (astrotypes.UserTimeObj, error) {

	var timeobj astrotypes.UserTimeObj

	_db := OpenUserDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	results, err := _db.Query("SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		panic(err)
	}

	if !results.NextResultSet() {
		return astrotypes.UserTimeObj{}, errors.New("ID not founded")
	}

	for results.Next() {
		var userobj astrotypes.UserTimeObj

		err = results.Scan(&userobj.Id, &userobj.Username, &userobj.TimeInSeconds, &userobj.Map)

		if err != nil {
			panic(err.Error())
		}

		timeobj = userobj
	}

	_, err = _db.Exec("UPDATE userRanking SET timeInSeconds=? WHERE id=?", tobj.TimeInSeconds, id)

	if err != nil {
		panic(err)
	}

	return timeobj, nil
}

func DeleteRanking(id string) (astrotypes.UserTimeObj, error) {

	var timeobj astrotypes.UserTimeObj

	_db := OpenUserDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	results, err := _db.Query("SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		panic(err)
	}

	if !results.NextResultSet() {
		return astrotypes.UserTimeObj{}, errors.New("ID not founded")
	}

	for results.Next() {
		var userobj astrotypes.UserTimeObj

		err = results.Scan(&userobj.Id, &userobj.Username, &userobj.TimeInSeconds, &userobj.Map)

		if err != nil {
			panic(err.Error())
		}

		timeobj = userobj
	}

	_, err = _db.Exec("DELETE FROM userRanking WHERE id=?", id)

	if err != nil {
		panic(err)
	}

	return timeobj, nil
}
