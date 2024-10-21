package main

import (
	"log"
	"os"

	"github.com/belovetech/gator.git/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	appState := &state{
		config: &config,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handleLogin)

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
