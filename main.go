package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/wodorek/blogator/internal/config"
	"github.com/wodorek/blogator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	appState := &state{cfg: &cfg}

	db, err := sql.Open("postgres", appState.cfg.DBURL)

	if err != nil {
		log.Fatal("error opening database connection")
	}

	dbQueries := database.New(db)
	appState.db = dbQueries

	cmds := commands{RegisteredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerGetAllFeeds)
	cmds.register("follow", handlerFeedFollow)
	cmds.register("following", handlerFeedFollowing)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(appState, command{Name: cmdName, Args: cmdArgs})

	if err != nil {
		log.Fatal(err)
	}

}
