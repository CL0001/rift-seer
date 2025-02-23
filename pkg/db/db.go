package db

import (
	"errors"
	"fmt"
	"github.com/CL0001/rift-seer/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to database failed: ", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("migration failed: ", err)
	}
}

func AddUser(user *models.User) error {
	if err := DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("email already in use")
		}
		return err
	}

	return nil
}

func FetchUser(email string) (models.User, error) {
	var user models.User

	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user does not exist")
		}
		return models.User{}, err
	}

	return user, nil
}