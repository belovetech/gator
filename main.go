package main

import (
	"fmt"
	"log"

	"github.com/belovetech/gator.git/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	fmt.Printf("Current User: %s\n", cfg.CurrentUserName)

	err = cfg.SetUser("new_user")
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	fmt.Println("User updated successfully.")
}
