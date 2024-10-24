# Gator

## RSS Aggregator

Gator is a simple RSS aggregator built with Go that allows users to manage and aggregate RSS feeds efficiently.

### Configuration

To set up the configuration for Gator, follow these steps:

1. Create the configuration file:

   ```bash
   touch ~/.gatorconfig.json
   ```

2. Edit the content of ~/.gatorconfig.json:
   ```json
   {
     "db_url": "postgres://<username>:@localhost:5432/gator?sslmode=disable"
   }
   ```
   Replace <username> with your PostgreSQL username.

### How to Use

Once the configuration is set up, you can use the following commands to interact with Gator:

```bash
    go run . reset       # Reset the application to its initial state

    go run . login       # Log in to your account

    go run . register    # Register a new user

    go run . users       # List all users

    go run . addfeed     # Add a new RSS feed

    go run . agg         # Aggregate the RSS feeds

    go run . feeds       # List all feeds

    go run . follow      # Follow a feed

    go run . following   # List the feeds you are following

    go run . unfollow    # Unfollow a feed
```

### Prerequisites

- Go (Golang) installed on your machine
- PostgreSQL database running locally
