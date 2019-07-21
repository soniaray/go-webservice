package models

type User struct {
	Id int
	Firstname string
	Lastname string
}

var (
	users []*User
	nextId = 1
)

func GetUsers() []*User {
	return  users
}

func AddUser(u User) (User, error) {
	u.Id = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}