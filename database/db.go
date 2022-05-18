package db

import (
	"log"
	"os"
	"path/filepath"
	"stumble/models"

	"github.com/joho/godotenv" // package used to read the .env file
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)
	log.Fatalln(err)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}

func fatal(err error) {
	panic("unimplemented")
}
