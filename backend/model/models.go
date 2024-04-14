package model

import (
	"gorm.io/gorm"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
}

type Pet struct {
	gorm.Model
	Name        string
	Type        string
	Description string
	Status      string
	ImageURL    string
}

type Service struct {
	gorm.Model
	Title       string
	Description string
}

type VolunteerOpportunity struct {
	gorm.Model
	Title              string
	Description        string
	AvailablePositions int
}

type Donation struct {
	gorm.Model
	UserID     uint
	User       User
	Amount     float64
	DonatedAt  time.Time
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
	return err
}

func (u *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}