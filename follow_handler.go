package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/belovetech/gator.git/internal/database"
	"github.com/google/uuid"
)

func handleFollow(state *state, cmd command) error {
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

	user, err := state.db.GetUser(ctx, state.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get the current user: %v", err)
	}
	_, err = state.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		if isUniqueConstraintViolation(err) {
			return fmt.Errorf("you are already following this feed")
		}
		return fmt.Errorf("unable to follow the feed: %v", err)
	}
	fmt.Println("The feed has been followed")
	fmt.Printf("%s\n%s\n", feed.Name, state.config.CurrentUserName)
	return nil
}

func isUniqueConstraintViolation(err error) bool {
	return strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint")
}
