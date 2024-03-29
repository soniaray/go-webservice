package models

import (
	"errors"
	"fmt"
)

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
	if u.Id != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.Id = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User,error){
	for _, u:= range users {
		if u.Id == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.Id == u.Id {
			users[i] = &u
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.Id)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}