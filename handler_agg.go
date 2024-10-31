package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"

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

	for _, item := range feed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    nextToFetch.ID,
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
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
