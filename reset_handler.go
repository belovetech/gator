package main

import (
	"context"
	"fmt"
)

func handleReset(state *state, cmd command) error {
	ctx := context.Background()

	err := state.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("unable to delete users: %v", err)
	}
	return nil
}
