package main

import (
	"context"
	"fmt"

	"github.com/belovetech/gator.git/internal/database"
)

func handleFollowing(state *state, cmd command, user database.User) error {
	ctx := context.Background()
	followedFeeds, err := state.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("unable to get followed feeds: %v", err)
	}

	for _, followedFeed := range followedFeeds {
		fmt.Printf("* %s\n", followedFeed.FeedName)
	}
	return nil
}
