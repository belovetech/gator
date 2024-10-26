package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/belovetech/gator.git/internal/database"
	"github.com/google/uuid"
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

	for _, feed := range feed.Channel.Items {
		pubDate, err := parsePubDate(feed.PubDate)
		if err != nil {
			return fmt.Errorf("unable to parse pub date: %v", err)
		}

		_, err = state.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			FeedID:      nextFeed.ID,
			Title:       feed.Title,
			Url:         feed.Link,
			Description: sql.NullString{String: feed.Description, Valid: true},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			PublishedAt: pubDate,
		})
		if err != nil {
			if isUniqueConstraintViolation(err) {
				continue
			}
			return fmt.Errorf("unable to create post: %v", err)
		}
		fmt.Printf("Created post: %s\n", feed.Title)
	}

	return nil
}

func parsePubDate(date string) (time.Time, error) {
	t, err := time.Parse(time.RFC1123, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
