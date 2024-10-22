package main

import (
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	username := cmd.args[0]
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to set the username")
	}
	fmt.Println("the user has been set")
	return nil
}
