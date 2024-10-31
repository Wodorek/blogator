package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wodorek/blogator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return errors.New("error fetching user, make sure you are logged in")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	createdFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: feedName, Url: feedUrl, UserID: user.ID})

	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: createdFeed.ID})

	if err != nil {
		return err
	}

	fmt.Println(createdFeed)

	return nil
}

func handlerGetAllFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetAllFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("Created by: %s\n", feed.Username)
		fmt.Println()
	}

	return nil
}
