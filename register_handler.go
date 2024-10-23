package main

import (
	"context"
	"fmt"
	"time"

	"github.com/belovetech/gator.git/internal/database"
	"github.com/google/uuid"
)

func handleRegister(state *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	name := cmd.args[0]
	if name == "" {
		return fmt.Errorf("user cannot be empty")
	}

	ctx := context.Background()

	userExists, err := state.db.GetUser(ctx, name)

	if err != nil && !isUserNotFound(err) {
		return fmt.Errorf("error checking user existence: %v", err)
	}

	if isUserAlreadyExists(userExists) {
		return fmt.Errorf("user already exists: %s", name)
	}

	user, err := state.db.CreateUser(ctx, database.CreateUserParams{
		Name:      name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("unable to create user: %v", err)
	}

	state.config.SetUser(name)

	fmt.Printf("the user has been created: %s\n", user.Name)
	return nil
}

func isUserNotFound(err error) bool {
	return err.Error() == "sql: no rows in result set"
}

func isUserAlreadyExists(user database.User) bool {
	return user.ID != uuid.Nil
}
