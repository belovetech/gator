package main

import (
	"context"
	"fmt"
	"time"

	"github.com/belovetech/gator.git/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("the addFeed handler expects two arguments: the name and feed URL")
	}
	name, feedURL := cmd.args[0], cmd.args[1]

	ctx := context.Background()
	createdFeed, err := state.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       feedURL,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		if isUniqueConstraintViolation(err) {
			return fmt.Errorf("feed already added")
		}
		return fmt.Errorf("unable to add the feed: %v", err)
	}

	_, err = state.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    createdFeed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("unable to follow the feed after creation: %v", err)
	}

	fmt.Println("The feed has been added")
	return nil
}
