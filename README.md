# Gator - RSS Feed Aggregator

Gator is a command-line RSS feed aggregator that helps you follow and manage your favorite RSS feeds.

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
go install github.com/fatkungfu/gator@latest
```

## Configuration

Manually create a config file in your home directory, `~/.gatorconfig.json`:

```json
{
  "db_url": "postgresql://username:@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username` with your PostgreSQL credentials.

## Usage

Here are some common commands:

```bash
# Register a new user
gator register <name>

# List all users
gator users

# Login as a user that already exists
gator login <name>

# Add a new RSS feed
gator addfeed <name> <url>

# List all feeds
gator feeds

# List feeds you're following
gator following

# Follow an existing feed that already exists in the database
gator follow https://blog.boot.dev/index.xml

# Unfollow a feed that already exists in the database
gator unfollow <url>

# Start the feed aggregator (runs continuously)
gator agg 1m

# Browse your posts
gator browse [limit]
gator browse 10   # shows 10 most recent posts
```

## Development

1. Clone the repository

```bash
git clone https://github.com/fatkungfu/gator.git
```

2. Install dependencies

```bash
go mod download
```

3. Create the database

- Enter the psql shell:

```bash
Mac: psql postgres
Linux: sudo -u postgres psql
```

- Create a new database. I called mine gator:
  `CREATE DATABASE gator;`

4. Run migrations

```bash
goose -dir sql/schema postgres "postgresql://username:@localhost:5432/gator" up
```

5. Build the project

```bash
go build
```
