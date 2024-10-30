package main

import (
	"context"
	"fmt"
	"log"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		log.Fatal("error fetching data")
	}

	fmt.Println(feed)

	return nil
}
