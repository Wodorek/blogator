package main

import (
	"fmt"
	"log"

	"github.com/wodorek/blogator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	fmt.Printf("Current config %v\n", cfg)

	err = cfg.SetUser("Wodorek")

	if err != nil {
		log.Fatalf("Error setting username: %v", err)
	}

	cfg, err = config.Read()

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	fmt.Printf("Config after save: %v\n", cfg)

}
