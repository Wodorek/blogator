package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wodorek/blogator/internal/database"
)

func handlerFeedFollow(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])

	if err != nil {
		return err
	}

	createdFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: currentUser.ID, FeedID: feed.ID})

	if err != nil {
		return err
	}

	fmt.Printf("User '%s' followed feed: '%s'\n", createdFollow.Username, createdFollow.FeedName)

	return nil
}

func handlerFeedFollowing(s *state, cmd command) error {

	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return err
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), currUser.ID)

	if err != nil {
		return err
	}

	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}
