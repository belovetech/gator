package main

import (
	"context"
	"fmt"
)

func handleUsers(state *state, cmd command) error {
	ctx := context.Background()

	users, err := state.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("unable to get users: %v", err)
	}
	for _, name := range users {

		currentUser := state.config.CurrentUserName

		if currentUser == name {
			fmt.Printf("* %s (current)\n", name)
			continue
		}

		fmt.Printf("* %s\n", name)
	}
	return nil
}
