package main

import (
	"context"
	"fmt"
)

func handleFeeds(state *state, cmd command) error {
	ctx := context.Background()

	feeds, err := state.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("unable to get feeds in the database")
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n%s\n", feed.Name, feed.UserName)
	}
	return nil
}
