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

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", config.DBUrl)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	dbQueries := database.New(db)

	appState := &state{
		db:     dbQueries,
		config: &config,
	}

	cmds := newCommandRegistry()

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Error: no command provided")
	}
	cmdName, cmdArgs := args[1], args[2:]

	registerCommands(cmds)
	runCommand(cmds, appState, cmdName, cmdArgs)
}
