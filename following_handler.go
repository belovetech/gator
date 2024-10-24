package main

import (
	"context"
	"fmt"
)

func handleFollowing(state *state, cmd command) error {

	username := state.config.CurrentUserName

	ctx := context.Background()
	user, err := state.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("unable to get user: %v", err)
	}
	followedFeeds, err := state.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("unable to get followed feeds: %v", err)
	}
	
	for _, followedFeed := range followedFeeds {
		fmt.Printf("* %s\n", followedFeed.FeedName)
	}
	return nil
}
