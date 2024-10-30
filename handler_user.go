package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wodorek/blogator/internal/database"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)

	if err != nil {
		return fmt.Errorf("no user %s found", username)
	}

	err = s.cfg.SetUser(username)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s has logged in\n", username)

	return nil

}

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: username})

	if err != nil {
		return fmt.Errorf("user %s already exists", username)
	}

	err = s.cfg.SetUser(username)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user %s created successfully\n", username)

	return nil
}

func handlerReset(s *state, cmd command) error {

	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.ResetTable(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func handlerGetUsers(s *state, cmd command) error {

	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetAllUsers(context.Background())

	if err != nil {
		return err
	}

	loggedUser := s.cfg.CurrentUserName

	for _, user := range users {

		isCurrent := ""

		if user.Name == loggedUser {
			isCurrent = "(current)"
		}

		fmt.Printf("* %s %s\n", user.Name, isCurrent)
	}

	return nil
}
