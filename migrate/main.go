package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jesperkha/w4d/config"
	"github.com/jesperkha/w4d/database"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to database...")
	db := database.New(config)

	log.Println("Attempting to migrate...")
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Println("Done")

	log.Println("Creating database dump...")

	dumpName := fmt.Sprintf("migrate/dumps/%s.sql", time.Now().Local().Format(time.RFC3339Nano))
	cmd := exec.Command("sqlite3", config.DbName, ".dump")

	file, err := os.Create(dumpName)
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}
	defer file.Close()

	cmd.Stdout = file

	err = cmd.Run()
	if err != nil {
		log.Fatal("Error running command: ", err)
	}

	log.Printf("SQL dump created successfully in %s\n", dumpName)
}
