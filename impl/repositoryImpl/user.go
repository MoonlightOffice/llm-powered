package repositoryImpl

import (
	"errors"
	"example/client/repository"
	"example/model/cterr"
	"example/model/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() repository.IUserRepository {
	return &userRepository{
		db: InitSQLite(),
	}
}

func (ur *userRepository) FindByUserId(userId string) (*user.User, error) {
	var u user.User
	result := ur.db.First(&u, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, cterr.ErrorNotExisted
		}
		return nil, result.Error
	}
	return &u, nil
}

func (ur *userRepository) Save(u user.User) error {
	result := ur.db.Save(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
