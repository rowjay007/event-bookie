package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rowjay007/event-bookie/config"
	"github.com/rowjay007/event-bookie/internal/router"
	"github.com/rowjay007/event-bookie/pkg/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize configuration
	conf := config.NewConfig()

	// Connect to the database
	db, err := database.NewDB(conf)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Convert *sql.DB to *gorm.DB
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error converting *sql.DB to *gorm.DB: %v", err)
	}

	// Apply migrations
	err = database.ApplyMigrations(db, conf)
	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	// Print a message to indicate successful migration
	fmt.Println("🔥 Migrations applied successfully 🌈🌈💥")

	// Initialize router with gormDB
	r := router.NewRouter(gormDB)

	// Start HTTP server
	port := conf.Port
	fmt.Printf("🌈 Server is running on port %s\n🔥", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
