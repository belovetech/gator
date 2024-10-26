package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/belovetech/gator.git/internal/database"
)

func handleBrowser(state *state, cmd command) error {
	limit := 2
	if len(cmd.args) > 0 {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("unable to convert limit (%v) to int", cmd.args[0])
		}
	}

	ctx := context.Background()
	posts, err := state.db.GetPostByUser(ctx, database.GetPostByUserParams{
		Name:  state.config.CurrentUserName,
		Limit: int32(limit),
	})
	if err != nil {
		fmt.Printf("Unable to get posts: %v", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\n", post.Title, post.Url)
		if post.Description.Valid {
			fmt.Println("Description: " + post.Description.String)
		} else {
			fmt.Println("No description available")
		}
		fmt.Println()
	}
	return nil
}
