package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jesperkha/w4d/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config := gorm.Config{}

	log.Println("Connecting to database...")
	db, err := gorm.Open(sqlite.Open("dev.db"), &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Attempting to migrate...")
	if err := db.AutoMigrate(&models.Recipe{}); err != nil {
		log.Fatal(err)
	}

	log.Println("Done")

	log.Println("Creating database dump...")

	dumpName := fmt.Sprintf("migrate/dumps/%s.sql", time.Now().Local().Format(time.RFC3339Nano))
	cmd := exec.Command("sqlite3", "dev.db", ".dump recipes")

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
