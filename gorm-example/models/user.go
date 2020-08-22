package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `User`
type User struct {
	gorm.Model
	Username string
	Email    string
}

func ConnectToDB() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		msg := fmt.Sprintf("Failed to connect to database! Reasone: err", err)
		panic(msg)
	}

	database.AutoMigrate(&User{})

	DB = database

}

func CreateUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include an ID, it must be set to 0")
	}
	results := DB.Create(&u)
	if results.Error != nil {
		panic("We hit an error creating a user")
	}
	return u, nil
}

func GetUserById(id int) (User, error) {
	var user User
	results := DB.Find(&user, id)
	if results.Error != nil {
		msg := fmt.Sprintf("Users with ID '%v' not found", id)
		return User{}, fmt.Errorf(msg)
	}
	return user, nil
}

// func UpdateUserEmail(id int, user User) (User, error) {
// 	for _, u := range users {
// 		if u.ID == id {
// 			u.Email = user.Email
// 			return *u, nil
// 		}
// 	}
// 	return User{}, fmt.Errorf("Unable to find user with ID '%v'", user.ID)
// }

// func DeleteUserById(id int) error {
// 	for i, u := range users {
// 		if u.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("User with id '%v' was not able to be deleted", id)
// }
