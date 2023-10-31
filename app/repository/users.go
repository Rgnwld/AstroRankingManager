package repository

import (
	AstroTypes "Astro/types"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, cred AstroTypes.DBCredentials) error
	GetUser(ctx context.Context, username string) (AstroTypes.DBCredentials, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, cred AstroTypes.DBCredentials) error {
	_, err := ur.db.ExecContext(ctx,
		"INSERT INTO userTable ( id, username, hashedpassword)  VALUES (?, ?, ?)",
		cred.Id, cred.Username, cred.HashedPassword)
	return err
}

func (ur *userRepository) GetUser(ctx context.Context, username string) (AstroTypes.DBCredentials, error) {
	var cred AstroTypes.DBCredentials

	results, err := ur.db.QueryContext(ctx, "SELECT * FROM  userTable WHERE username=?", username)
	if err != nil {
		return cred, err
	}

	defer results.Close()

	for results.Next() {
		err = results.Scan(&cred.Id, &cred.Username, &cred.HashedPassword)
		if err != nil {
			return cred, err
		}
	}
	return cred, nil
}
