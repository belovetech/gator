package main

import (
	"fmt"

	"github.com/belovetech/gator.git/internal/config"
)

func main() {
	fmt.Println("Gator: A simple CLI task manager")

	// Read the config file
	config, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	} else {
		fmt.Println("Config file read successfully")
		fmt.Println("DB URL:", config.DBUrl)
		fmt.Println("Current User Name:", config.CurrentUserName)
	}

	config.SetUser("John Doe")

}
