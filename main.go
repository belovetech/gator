package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/belovetech/gator.git/internal/config"
	"github.com/belovetech/gator.git/internal/database"
	_ "github.com/lib/pq"
)

func loadConfig() (config.Config, error) {
	return config.Read()
}

func initializeState() (*database.Queries, *config.Config, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, nil, err
	}

	db, err := sql.Open("postgres", config.DBUrl)
	if err != nil {
		return nil, nil, err
	}

	return database.New(db), &config, nil

}

func parseCommandArgs() (string, []string) {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Error: no command provided")
	}
	return args[1], args[2:]
}

func main() {
	dbQueries, config, err := initializeState()
	if err != nil {
		log.Fatalf("failed to initialize state: %v", err)
	}

	appState := &state{
		db:     dbQueries,
		config: config,
	}

	cmdName, cmdArgs := parseCommandArgs()
	cmds := newCommandRegistry()
	registerCommands(cmds)
	runCommand(cmds, appState, cmdName, cmdArgs)
}
