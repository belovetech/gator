package main

import (
	"fmt"

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

// func handleLogin(s *state, cmd command) error {
// 	if len(cmd.args) < 1 {
// 		return fmt.Errorf("the login handler expects a single argument, the username")
// 	}

// 	username := cmd.args[0]
// 	if username == "" {
// 		return fmt.Errorf("username cannot be empty")
// 	}

// 	err := s.config.SetUser(username)
// 	if err != nil {
// 		return fmt.Errorf("unable to set the username")
// 	}
// 	fmt.Println("the user has been set")
// 	return nil
// }
