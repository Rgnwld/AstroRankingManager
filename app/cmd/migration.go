package cmd

import (
	DBConn "Astro/database"
	"context"
	"errors"
	"fmt"
	"github.com/pressly/goose/v3"
	"os"
	"sync"
	"time"
)

var errInvalidMigrationAction = errors.New("invalid migration action")
var migrationDbOnce sync.Once

func (app *app) migrationUp() error {
	// TODO: figure out which timeout is better to wait
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	migrationDb := DBConn.InitializeDB(
		app.config.migrationDbOpts.user,
		app.config.migrationDbOpts.pass,
		app.config.migrationDbOpts.addr,
		app.config.migrationDbOpts.dbName,
		&migrationDbOnce)

	goose.SetBaseFS(os.DirFS("./"))

	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("migrations: fail to select dialect: %w", err)
	}

	// default behavior is up migrations
	switch app.config.migrationDbOpts.action {
	case "reset":
		fmt.Println("reset all migrations")
		return goose.ResetContext(ctx, migrationDb, "migrations")
	case "up":
		fmt.Println("setup migrations")
		return goose.UpContext(ctx, migrationDb, "migrations")
	default:
		return errInvalidMigrationAction
	}
}
