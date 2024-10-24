package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/belovetech/gator.git/internal/database"
)

func scrapeFeeds(state *state) error {
	ctx := context.Background()

	nextFeed, err := state.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("unable to get the next feed to fetch: %v", err)
	}

	now := time.Now()
	err = state.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		UpdatedAt:     now,
		LastFetchedAt: sql.NullTime{Time: now, Valid: true},
		ID:            nextFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("unable to mark the feed as fetched: %v", err)
	}

	feed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %v", err)
	}

	fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	fmt.Println("--------------------")

	fmt.Println("Feed Items:")
	fmt.Println("--------------------")
	for _, feed := range feed.Channel.Items {
		fmt.Printf("%s\n", feed.Title)
	}

	fmt.Println("--------------------")

	return nil
}
