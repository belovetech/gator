package main

import (
	"context"
	"fmt"
	"time"

	"github.com/belovetech/gator.git/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(state *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("the addFeed handler expects two arguments: the name and feed URL")
	}
	name, feedURL := cmd.args[0], cmd.args[1]

	currentUserName := state.config.CurrentUserName
	ctx := context.Background()
	user, err := state.db.GetUser(ctx, currentUserName)
	if err != nil {
		return fmt.Errorf("unable to get the current user: %v", err)
	}

	feedExists, err := state.db.GetFeed(ctx, feedURL)
	if err != nil && !isFeedNotFound(err) {
		return fmt.Errorf("error checking feed existence: %v", err)
	}

	if isFeedAlreadyExists(feedExists) {
		feed, err := fetchFeed(ctx, feedURL)
		if err != nil {
			return fmt.Errorf("unable to fetch feed: %v", err)
		}
		fmt.Printf("Feed: %v\n", feed)
		return nil
	}

	_, err = state.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       feedURL,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("unable to add the feed: %v", err)
	}

	feed, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("unable to fetch feed after creation: %v", err)
	}
	fmt.Println("The feed has been added")
	fmt.Printf("Feed: %v\n", feed)
	return nil
}

func isFeedNotFound(err error) bool {
	return err.Error() == "sql: no rows in result set"
}

func isFeedAlreadyExists(feed database.Feed) bool {
	return feed.ID == uuid.Nil
}
