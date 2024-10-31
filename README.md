# Blogator

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `blogator` with:

```bash
go install ...
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
blogator register <name>
```

Add a feed:

```bash
blogator addfeed <url>
```

Start the aggregator:

```bash
blogator agg 30s
```

View the posts:

```bash
blogator browse [limit]
```

There are a few other commands you'll need as well:

- `blogator login <name>` - Log in as a user that already exists
- `blogator users` - List all users
- `blogator feeds` - List all feeds
- `blogator follow <url>` - Follow a feed that already exists in the database
- `blogator unfollow <url>` - Unfollow a feed that already exists in the database
