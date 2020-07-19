package models

type User struct {
	ID       int
	Username string
	Email    string
}

var (
	users  []*User //slice of users with pointer to user struct
	nextID = 1
)
