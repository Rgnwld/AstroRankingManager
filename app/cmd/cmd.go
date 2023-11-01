package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type dbOpts struct {
	user   string
	pass   string
	addr   string
	dbName string
}

type config struct {
	rankingDbOpts   dbOpts
	usersDbOpts     dbOpts
	migrationDbOpts struct {
		dbOpts
		action string
	}
}

type cmdSet struct {
	fs  *flag.FlagSet
	run func() error
}

type app struct {
	// TODO: setup logger
	config      config
	subCommands map[string]*cmdSet
}

var errNotValidCommand = errors.New("Add the subcommand 'api' or 'migration' ")

// Execute parse and execute commands/subcommands
// Ref: https://paulgorman.org/technical/golang-flag.txt.html
func Execute() error {

	if err := LoadDotEnvVariables(); err != nil {
		fmt.Println(err)
	}

	var app = &app{subCommands: map[string]*cmdSet{}}

	// Setup api command
	api := flag.NewFlagSet("api", flag.ExitOnError)
	app.subCommands["api"] = &cmdSet{fs: api, run: app.serveApi}

	api.StringVar(&app.config.rankingDbOpts.user, "rankingDbUser", os.Getenv("DB_RANKING_USER"), "ranking Db option 'user'")
	api.StringVar(&app.config.rankingDbOpts.pass, "rankingDbPass", os.Getenv("DB_RANKING_PASS"), "ranking Db option 'pass'")
	api.StringVar(&app.config.rankingDbOpts.addr, "rankingDbAddr", os.Getenv("DB_RANKING_ADDR"), "ranking Db option 'addr'")
	api.StringVar(&app.config.rankingDbOpts.dbName, "rankingDbName", os.Getenv("DB_RANKING_NAME"), "ranking Db option 'dbName'")

	api.StringVar(&app.config.usersDbOpts.user, "usersDbUser", os.Getenv("DB_USERS_USER"), "users Db option 'user'")
	api.StringVar(&app.config.usersDbOpts.pass, "usersDbPass", os.Getenv("DB_USERS_PASS"), "users Db option 'pass'")
	api.StringVar(&app.config.usersDbOpts.addr, "usersDbAddr", os.Getenv("DB_USERS_ADDR"), "users Db option 'addr'")
	api.StringVar(&app.config.usersDbOpts.dbName, "usersDbName", os.Getenv("DB_USERS_NAME"), "users Db option 'dbName'")

	// Setup migration command
	migration := flag.NewFlagSet("api", flag.ExitOnError)
	app.subCommands["migration"] = &cmdSet{fs: migration, run: app.migrationUp}

	migration.StringVar(&app.config.migrationDbOpts.user, "dbUser", os.Getenv("DB_MIGRATION_USER"), "migration Db option 'user'")
	migration.StringVar(&app.config.migrationDbOpts.pass, "dbPass", os.Getenv("DB_MIGRATION_PASS"), "migration Db option 'pass'")
	migration.StringVar(&app.config.migrationDbOpts.addr, "dbAddr", os.Getenv("DB_MIGRATION_ADDR"), "migration Db option 'addr'")
	migration.StringVar(&app.config.migrationDbOpts.dbName, "dbName", os.Getenv("DB_MIGRATION_NAME"), "migration Db option 'dbName'")
	migration.StringVar(&app.config.migrationDbOpts.action, "action", "up", "migration action 'action'")

	// Switch on subcommands, then apply the desired set of flags.
	if len(os.Args) < 2 {
		return errNotValidCommand
	}

	cmdName := os.Args[1]
	cmd, ok := app.subCommands[cmdName]
	if !ok {
		return errNotValidCommand
	}

	if err := cmd.fs.Parse(os.Args[2:]); err != nil {
		return err
	}

	return cmd.run()
}
