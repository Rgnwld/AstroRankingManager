package DBConn

import (
	astrotypes "Astro/types"
	"database/sql"
	"errors"
)

var (
	errIdNotFound = errors.New("ID not founded")
)

type RankingRepository struct {
	db *sql.DB
}

func NewRankingRepository(db *sql.DB) *RankingRepository {
	return &RankingRepository{
		db: db,
	}
}

func (rr *RankingRepository) AddRanking(tobj astrotypes.UserTimeObj) error {
	if _, err := rr.db.Exec("USE AstroRankings"); err != nil {
		return err
	}

	_, err := rr.db.Exec("INSERT INTO userRanking ( id, userId, timeInSeconds, mapId)  VALUES (?, ?, ?, ?);",
		tobj.Id, tobj.UserId, tobj.TimeInSeconds, tobj.MapId)
	if err != nil {
		return err
	}

	return nil
}

func (rr *RankingRepository) GetRankings() ([]astrotypes.UserTimeObj, error) {
	results, err := rr.db.Query("SELECT * FROM userRanking ORDER BY timeInSeconds ASC")
	if err != nil {
		return []astrotypes.UserTimeObj{}, err
	}

	defer func() {
		_ = results.Close()
	}()

	var times []astrotypes.UserTimeObj

	for results.Next() {
		var userTimes astrotypes.UserTimeObj

		if err = results.Scan(
			&userTimes.Id,
			&userTimes.UserId,
			&userTimes.TimeInSeconds,
			&userTimes.MapId); err != nil {
			return []astrotypes.UserTimeObj{}, err
		}

		times = append(times, userTimes)
	}

	return times, nil
}

func (rr *RankingRepository) GetPlayerAllRanking(playerId string) ([]astrotypes.UserTimeObj, error) {
	results, err := rr.db.Query("SELECT * FROM userRanking WHERE userId=?", playerId)
	if err != nil {
		return []astrotypes.UserTimeObj{}, err
	}

	defer func() {
		_ = results.Close()
	}()

	var times []astrotypes.UserTimeObj

	for results.Next() {
		var userTimes astrotypes.UserTimeObj

		err = results.Scan(&userTimes.Id, &userTimes.UserId, &userTimes.TimeInSeconds, &userTimes.MapId)

		if err != nil {
			return []astrotypes.UserTimeObj{}, err
		}

		times = append(times, userTimes)
	}

	return times, nil
}

func (rr *RankingRepository) GetRankingByMap(mapId string) ([]astrotypes.UserTimeObj, error) {
	results, err := rr.db.Query("SELECT * FROM userRanking WHERE mapId=?", mapId)
	if err != nil {
		panic(err)
	}

	var times []astrotypes.UserTimeObj

	for results.Next() {
		var userTimes astrotypes.UserTimeObj

		err = results.Scan(&userTimes.Id, &userTimes.UserId, &userTimes.TimeInSeconds, &userTimes.MapId)

		if err != nil {
			return []astrotypes.UserTimeObj{}, err
		}

		times = append(times, userTimes)
	}

	return times, nil
}

func (rr *RankingRepository) PatchRankingTime(id string, tobj astrotypes.TimeObj) (astrotypes.UserTimeObj, error) {
	var timeobj astrotypes.UserTimeObj

	results, err := rr.db.Query("SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		return timeobj, err
	}

	if !results.NextResultSet() {
		return astrotypes.UserTimeObj{}, errIdNotFound
	}
	defer func() {
		_ = results.Close()
	}()

	for results.Next() {
		var userobj astrotypes.UserTimeObj

		err = results.Scan(&userobj.Id, &userobj.UserId, &userobj.TimeInSeconds, &userobj.MapId)

		if err != nil {
			return astrotypes.UserTimeObj{}, err
		}

		timeobj = userobj
	}

	_, err = rr.db.Exec("UPDATE userRanking SET timeInSeconds=? WHERE id=?", tobj.TimeInSeconds, id)
	if err != nil {
		return astrotypes.UserTimeObj{}, err
	}

	return timeobj, nil
}

func (rr *RankingRepository) DeleteRanking(id string) (astrotypes.UserTimeObj, error) {

	var timeobj astrotypes.UserTimeObj

	_, err := rr.db.Exec("USE AstroRankings")
	if err != nil {
		return timeobj, err
	}

	results, err := rr.db.Query("SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		return timeobj, err
	}

	defer func() {
		_ = results.Close()
	}()

	if !results.NextResultSet() {
		return astrotypes.UserTimeObj{}, errors.New("ID not founded")
	}

	for results.Next() {
		var userobj astrotypes.UserTimeObj

		err = results.Scan(&userobj.Id, &userobj.UserId, &userobj.TimeInSeconds, &userobj.MapId)
		if err != nil {
			return timeobj, err
		}

		timeobj = userobj
	}

	_, err = rr.db.Exec("DELETE FROM userRanking WHERE id=?", id)
	if err != nil {
		return timeobj, err
	}
	return timeobj, nil
}
