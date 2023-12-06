package repository

import (
	astrotypes "Astro/types"
	"context"
	"database/sql"
	"errors"
)

var (
	errIdNotFound = errors.New("ID not founded")
)

type RankingRepository interface {
	AddRanking(ctx context.Context, u astrotypes.UserTimeObj) error
	GetRankings(ctx context.Context) ([]astrotypes.UserTimeObj, error)
	GetPlayerAllRanking(ctx context.Context, playerId string) ([]astrotypes.UserTimeObj, error)
	GetRankingByMap(ctx context.Context, mapId string) ([]astrotypes.UserNameTimeObj, error)
	PatchRankingTime(ctx context.Context, id string, tobj astrotypes.TimeObj) (astrotypes.UserTimeObj, error)
	DeleteRanking(ctx context.Context, id string) (astrotypes.UserTimeObj, error)
}

type rankingRepository struct {
	db *sql.DB
}

func NewRankingRepository(db *sql.DB) RankingRepository {
	return &rankingRepository{
		db: db,
	}
}

func (rr *rankingRepository) AddRanking(ctx context.Context, tobj astrotypes.UserTimeObj) error {
	_, err := rr.db.ExecContext(ctx, "INSERT INTO userRanking ( id, userId, timeInMiliSeconds, mapId)  VALUES (?, ?, ?, ?);",
		tobj.Id, tobj.UserId, tobj.TimeInMiliSeconds, tobj.MapId)
	return err
}

func (rr *rankingRepository) GetRankings(ctx context.Context) ([]astrotypes.UserTimeObj, error) {
	var times []astrotypes.UserTimeObj

	results, err := rr.db.QueryContext(ctx, "SELECT * FROM userRanking ORDER BY timeInMiliSeconds ASC")
	if err != nil {
		return times, err
	}

	defer func() {
		_ = results.Close()
	}()

	for results.Next() {
		var userTimes astrotypes.UserTimeObj

		if err = results.Scan(
			&userTimes.Id,
			&userTimes.UserId,
			&userTimes.TimeInMiliSeconds,
			&userTimes.MapId); err != nil {
			return times, err
		}

		times = append(times, userTimes)
	}

	return times, nil
}

func (rr *rankingRepository) GetPlayerAllRanking(ctx context.Context, playerId string) ([]astrotypes.UserTimeObj, error) {
	var times []astrotypes.UserTimeObj

	results, err := rr.db.QueryContext(ctx, "SELECT * FROM userRanking WHERE userId=?", playerId)
	if err != nil {
		return times, err
	}

	defer func() {
		_ = results.Close()
	}()

	for results.Next() {
		var userTimes astrotypes.UserTimeObj

		err = results.Scan(&userTimes.Id, &userTimes.UserId, &userTimes.TimeInMiliSeconds, &userTimes.MapId)
		if err != nil {
			return times, err
		}

		times = append(times, userTimes)
	}

	return times, nil
}

func (rr *rankingRepository) GetRankingByMap(ctx context.Context, mapId string) ([]astrotypes.UserNameTimeObj, error) {
	var times []astrotypes.UserNameTimeObj
	results, err := rr.db.QueryContext(ctx, "SELECT  r.id, u.username AS username, r.timeInMiliSeconds , r.mapId FROM AstroRankings.userTable AS u JOIN AstroRankings.userRanking AS r ON r.userId=u.id WHERE mapId=?", mapId)
	if err != nil {
		return times, err
	}
	defer results.Close()

	for results.Next() {
		var userTimes astrotypes.UserNameTimeObj

		err = results.Scan(&userTimes.Id, &userTimes.Username, &userTimes.TimeInMiliSeconds, &userTimes.MapId)

		if err != nil {
			return times, err
		}
		times = append(times, userTimes)
	}
	return times, nil
}

func (rr *rankingRepository) PatchRankingTime(ctx context.Context, id string, tobj astrotypes.TimeObj) (astrotypes.UserTimeObj, error) {
	var timeobj astrotypes.UserTimeObj

	results, err := rr.db.QueryContext(ctx, "SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		return timeobj, err
	}

	if !results.NextResultSet() {
		return timeobj, errIdNotFound
	}
	defer func() {
		_ = results.Close()
	}()

	for results.Next() {
		if err := results.Scan(&timeobj.Id, &timeobj.UserId, &timeobj.TimeInMiliSeconds, &timeobj.MapId); err != nil {
			return timeobj, err
		}
	}

	_, err = rr.db.ExecContext(ctx, "UPDATE userRanking SET timeInMiliSeconds=? WHERE id=?", tobj.TimeInMiliSeconds, id)
	if err != nil {
		return timeobj, err
	}

	return timeobj, nil
}

func (rr *rankingRepository) DeleteRanking(ctx context.Context, id string) (astrotypes.UserTimeObj, error) {

	var timeobj astrotypes.UserTimeObj

	results, err := rr.db.QueryContext(ctx, "SELECT * FROM userRanking WHERE id=?", id)
	if err != nil {
		return timeobj, err
	}

	defer func() {
		_ = results.Close()
	}()

	if !results.NextResultSet() {
		return timeobj, errIdNotFound
	}

	for results.Next() {
		err = results.Scan(&timeobj.Id, &timeobj.UserId, &timeobj.TimeInMiliSeconds, &timeobj.MapId)
		if err != nil {
			return timeobj, err
		}
	}

	_, err = rr.db.ExecContext(ctx, "DELETE FROM userRanking WHERE id=?", id)
	if err != nil {
		return timeobj, err
	}
	return timeobj, nil
}
