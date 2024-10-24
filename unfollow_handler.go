package main

import (
	"context"
	"fmt"

	"github.com/belovetech/gator.git/internal/database"
)

func handleUnfollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("the follow handler expects a single argument, the feedURL")
	}

	feedURL := cmd.args[0]
	if feedURL == "" {
		return fmt.Errorf("feedURL cannot be empty")
	}

	ctx := context.Background()
	feed, err := state.db.GetFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("unable to get the feed: %v", err)
	}

	err = state.db.DeleteFeedFollows(ctx, database.DeleteFeedFollowsParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})

	if err != nil {
		return fmt.Errorf("unable to unfollow the feed: %v", err)
	}
	fmt.Println("The feed has been unfollowed")
	return nil
}
