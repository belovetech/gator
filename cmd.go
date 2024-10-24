package main

import (
	"fmt"
	"log"

	"github.com/belovetech/gator.git/internal/config"
	"github.com/belovetech/gator.git/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	if _, exists := c.handlers[name]; exists {
		panic(fmt.Sprintf("command %s already exists", name))
	}
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func registerCommands(cmds commands) {
	commandList := []struct {
		name    string
		handler func(*state, command) error
	}{
		{"login", handleLogin},
		{"register", handleRegister},
		{"reset", handleReset},
		{"users", handleUsers},
		{"agg", handlerAgg},
		{"addfeed", handleAddFeed},
		{"feeds", handleFeeds},
	}

	for _, cmd := range commandList {
		cmds.register(cmd.name, cmd.handler)
	}

}

func runCommand(cmds commands, appState *state, cmdName string, cmdArgs []string) error {
	err := cmds.run(appState, command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return nil
}
