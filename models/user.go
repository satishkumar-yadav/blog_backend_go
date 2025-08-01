package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID    string `json:"user_id" gorm:"primaryKey;size:36" `
	FirstName string `json:"first_name"  gorm:"size:50"`
	LastName  string `json:"last_name"  gorm:"size:50"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	Phone     string `json:"phone"  gorm:"size:16"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//Blogs    []Blog
}

// type User struct {
// 	Password  []byte `json: "-"` // "password"

// }

func (user *User) SetPassword(password string) error {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// user.Password = hashedPassword
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

/*

bcrypt.GenerateFromPassword - returns hashed password as []byte

*/
