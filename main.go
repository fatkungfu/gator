package main

import (
	"log"
	"os"

	"github.com/fatkungfu/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	// Read initial config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Create state and commands
	programState := &state{cfg: &cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// Register commands
	cmds.register("login", handlerLogin)

	// Ensure we have at least one command
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	// Create and run command
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
