package DBConn

import (
	AstroTypes "Astro/types"
	"database/sql"
)

func OpenUserDBConnection() *sql.DB {

	db, err := sql.Open("mysql", Cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	return db
}

func CreateUser(cred AstroTypes.DBCredentials) {
	_db := OpenUserDBConnection()
	defer _db.Close()

	_, err := _db.Exec("USE AstroRankings")
	if err != nil {
		panic(err)
	}

	_, err = _db.Exec("INSERT INTO userTable ( id, username, hashedpassword)  VALUES (?, ?, ?);",
		cred.Id, cred.Username, cred.HashedPassword)

	if err != nil {
		panic(err)
	}
}

func GetUser(username string) (AstroTypes.DBCredentials, error) {
	_db := OpenUserDBConnection()
	defer _db.Close()

	results, err := _db.Query("SELECT * FROM  userTable WHERE username=?", username)
	if err != nil {
		panic(err)
	}

	var cred AstroTypes.DBCredentials
	for results.Next() {

		err = results.Scan(&cred.Id, &cred.Username, &cred.HashedPassword)
		if err != nil {
			return AstroTypes.DBCredentials{}, err
		}

	}

	return cred, nil
}
