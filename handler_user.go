package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	fmt.Println(len(cmd.Args))

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	fmt.Println(username)

	err := s.cfg.SetUser(username)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s has logged in\n", username)

	return nil

}
