package main

import (
	"context"
	"fmt"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(state *state, cmd command) error {
	ctx := context.Background()

	feeds, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %v", err)
	}

	fmt.Printf("Feed: %v\n", feeds)

	return nil
}
