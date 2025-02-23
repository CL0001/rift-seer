package utils

import (
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func LoadEnv() {
	//".env"
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load environment variables: ", err)
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(enteredPassword, storedHash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(enteredPassword)); err != nil {
		return err
	}
	return nil
}
