package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wodorek/blogator/internal/database"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])

	if err != nil {
		return err
	}

	createdFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})

	if err != nil {
		return err
	}

	fmt.Printf("User '%s' followed feed: '%s'\n", createdFollow.Username, createdFollow.FeedName)

	return nil
}

func handlerFeedFollowing(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func handlerFeedUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{Url: cmd.Args[0], UserID: user.ID})

	if err != nil {
		return err
	}

	return nil

}
