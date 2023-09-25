package models

import (
	"jwt-go/database"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primary_key"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) CreateUser() error {
	result := database.GlobalDB.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))

	if err != nil {
		return err
	}

	return nil
}
