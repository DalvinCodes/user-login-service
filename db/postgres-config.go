package db

import (
	"fmt"
	"log"
	"os"
	"user-login-service/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupPostgresDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(GetPostgresDbDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Exec("DROP TABLE Users, Addresses")
	
	if err := db.AutoMigrate(
		// models to migrate in db
		&domain.User{},
		&domain.Address{},

		); err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	return db, nil
}

func GetPostgresDbDSN() string {
	if err := godotenv.Load(); err != nil {
		return err.Error()
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"))
}
