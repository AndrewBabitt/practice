package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
}

var (
	users  []*User //slice of users with pointer to user struct
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include an ID, it must be set to 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("Users with ID '%v' not found", id)
}

func UpdateUserEmail(id int, user User) (User, error) {
	for _, u := range users {
		if u.ID == id {
			u.Email = user.Email
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("Unable to find user with ID '%v'", user.ID)
}

func DeleteUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id '%v' was not able to be deleted", id)
}
