package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/belovetech/gator.git/internal/config"
	"github.com/belovetech/gator.git/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Read()
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

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Error: no command provided")
	}

	cmdName, cmdArgs := args[1], args[2:]
	err = cmds.run(appState, command{name: cmdName, args: cmdArgs})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
