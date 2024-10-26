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

Clone the repository:

```bash
   git clone https://github.com/belovetech/gator.git
```

Change directory to the cloned repository:

install goose CLI:

```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
   goose -v
```

Install the dependencies:

```bash
   go mod tidy
```

Run the migrations:

```bash
   make migrate-up
```

Build the application:

```bash
   go build -o gator
```

Once the configuration is set up, you can use the following commands to interact with Gator:

```bash
   ./gator reset                          # Reset the application to its initial state

   ./gator login <username>               # Log in to your account

   ./gator register  <username>           # Register a new user

   ./gator users                          # List all users

   ./gator addfeed <name> <URL>           # Add a new RSS feed

   ./gator agg <time_between_reqs>        # Aggregate the RSS feeds

   ./gator feeds                          # List all feeds

   ./gator follow  <URL>                  # Follow a feed

   ./gator unfollow  <URL>                # Unfollow a feed

   ./gator following                      # List the feeds you are following

   ./gator browse <LIMIT>                 # Browse the posts

```

### Prerequisites

- Go (Golang) installed on your machine
- PostgreSQL database running locally
