package repository

import "example/model/user"

type IUserRepository interface {
	FindByUserId(userId string) (*user.User, error)
	Save(userId user.User) error
}
