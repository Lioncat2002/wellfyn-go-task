package models

import (
	"main/models/database"
	"main/utils/token"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Name     string `gorm:"size:255;" json:"name"`
	Imgurl   string `gorm:"size:255;" json:"imgurl"`
	Phno     string `gorm:"size:255;" json:"phno"`
	Resume   string `gorm:"size:255;" json:"resume"`
}

// VerifyPassword used for verifying password
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// LoginCheck used for checking if the user exists in the database
func LoginCheck(email string, password string) (string, error) {

	var err error

	u := User{}

	err = database.DB.Model(User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.Email) //Generate new JWT token for the user

	if err != nil {
		return "", err
	}

	return token, nil

}

// UpdateUser used for updating user
func UpdateUser(u *User) (*User, error) {
	err := database.DB.Save(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// GetUserByEmail used for getting user by email since it's unique
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// Save used for saving user
func (u *User) Save() (*User, error) {
	err := database.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Encrypt used for encrypting password
func (u *User) Encrypt() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}
