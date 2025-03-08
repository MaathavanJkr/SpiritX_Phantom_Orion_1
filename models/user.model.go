// models/user.go
package models

import (
	"go-orm-template/db"
)

type User struct {
	GormModel
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Approved bool   `json:"approved"`
}
type UserWithPassword struct {
	User
	Password string `json:"password"`
}

// AddUser creates a new user record in the database
func AddUser(user *User) error {
	result := db.ORM.Create(&user)
	return result.Error
}

// GetUserByID retrieves a user record from the database by ID
func GetUserByID(id string) (*User, error) {
	var user *User
	result := db.ORM.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserByUsername retrieves a user record from the database by ID
func GetUserByUsername(username string) (*User, error) {
	user := &User{Username: username}
	result := db.ORM.Where(user).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetAllUsers() ([]*User, error) {
	var users []*User

	result := db.ORM.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// UpdateUserByID updates an existing user record in the database
func UpdateUserByID(user *User) error {
	result := db.ORM.Save(&user)
	return result.Error
}

// DeleteUserByID deletes a user record from the database by ID
func DeleteUserByID(id string) error {
	var user *User
	result := db.ORM.Delete(&user, id)
	return result.Error
}
