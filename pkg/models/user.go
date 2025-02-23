package models

import (
	"github.com/CL0001/rift-seer/pkg/utils"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username     string    `gorm:"not null"`
	SummonerName string    `gorm:"not null"`
	Email        string    `gorm:"unique;not null"`
	Password     string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func NewUser(username, summonerName, email, password string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return &User{}, err
	}

	user := &User{
		ID:           uuid.New(),
		Username:     username,
		SummonerName: summonerName,
		Email:        email,
		Password:     hashedPassword,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return user, nil
}
