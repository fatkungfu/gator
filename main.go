package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/fatkungfu/gator/internal/config"
	"github.com/fatkungfu/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	// Read initial config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	// Create state and commands
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// Register commands
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

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
