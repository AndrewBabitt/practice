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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
