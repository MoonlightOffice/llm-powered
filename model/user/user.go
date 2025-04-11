package user

import (
	"example/util"
)

type User struct {
	UserId string
	Name   string
	Age    int
}

func NewUser(name string, age int) User {
	return User{
		UserId: util.GenId("user"),
		Name:   name,
		Age:    age,
	}
}
