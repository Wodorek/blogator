package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) == 0 {
		return errors.New("no username provided")
	}

	username := cmd.Args[0]

	err := s.cfg.SetUser(username)

	if err != nil {
		return err
	}

	fmt.Printf("User %s has logged in\n", username)

	return nil

}
