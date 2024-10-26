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

func initializeState() (*state, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", config.DBUrl)
	if err != nil {
		return nil, err
	}

	appState := &state{
		db:     database.New(db),
		config: &config,
	}

	return appState, nil

}

func parseCommandArgs() (string, []string) {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Error: no command provided")
	}
	return args[1], args[2:]
}

func main() {
	appState, err := initializeState()
	if err != nil {
		log.Fatalf("failed to initialize state: %v", err)
	}
	cmdName, cmdArgs := parseCommandArgs()
	cmds := newCommandRegistry()
	registerCommands(cmds)
	runCommand(cmds, appState, cmdName, cmdArgs)
}
