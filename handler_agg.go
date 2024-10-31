package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/wodorek/blogator/internal/database"
)

func scrapeFeeds(s *state) error {

	nextToFetch, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}, ID: nextToFetch.ID})

	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), nextToFetch.Url)

	if err != nil {
		return err
	}

	fmt.Printf("articles from %s feed:\n", nextToFetch.Name)
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}

	fmt.Println()

	return nil
}

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time string>", cmd.Name)
	}

	timeTick, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeTick)

	fmt.Printf("Fetching feeds every %s \n", cmd.Args[0])

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
