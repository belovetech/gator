package main

import (
	"fmt"
	"time"
)

func handlerAgg(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: scrapefeeds <time_between_reqs>")
	}

	timeBetweenRequests := cmd.args[0]
	if timeBetweenRequests == "" {
		return fmt.Errorf("time between requests cannot be empty")
	}

	fmt.Printf("Collecting feeds every %s\n\n", timeBetweenRequests)

	// Parse the duration string into a time.Duration
	timeBetweenReq, err := time.ParseDuration(timeBetweenRequests)
	if err != nil {
		return fmt.Errorf("invalid time format for time between requests: %v", err)
	}

	ticker := time.NewTicker(timeBetweenReq)

	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
}
