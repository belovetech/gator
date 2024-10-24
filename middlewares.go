package main

import (
	"context"
	"fmt"

	"github.com/belovetech/gator.git/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Fetch the user from the state (database or config)
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to get logged-in user: %w", err)
		}

		// Call the handler with the user information
		return handler(s, cmd, user)
	}
}
