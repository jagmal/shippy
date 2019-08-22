package main

import (
	pbUser "github.com/jagmal/shippy/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll() ([]*pbUser.User, error)
	Get(id string) (*pbUser.User, error)
	Create(user *pbUser.User) error
	GetByEmailAndPassword(user *pbUser.User) (*pbUser.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pbUser.User, error) {
	var users []*pbUser.User

	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*pbUser.User, error) {
	var user *pbUser.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pbUser.User) (*pbUser.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pbUser.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
