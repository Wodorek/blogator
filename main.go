package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wodorek/blogator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	appState := &state{cfg: &cfg}
	cmds := commands{RegisteredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(appState, command{Name: cmdName, Args: cmdArgs})

	if err != nil {
		log.Fatal(err)
	}

}
