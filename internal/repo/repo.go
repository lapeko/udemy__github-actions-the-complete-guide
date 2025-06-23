package repo

import (
	"github.com/lapeko/udemy__github-actions-the-complete-guide/internal/model"
)

type repoInstance struct{}

type Repo interface {
	CreateUser(user *model.User) error
	GetUsers(size uint, page uint) ([]*model.User, error)
	GetUser(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

func New() Repo {
	return &repoInstance{}
}

func (r *repoInstance) CreateUser(user *model.User) error {
	return nil
}

func (r *repoInstance) GetUsers(size uint, page uint) ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *repoInstance) GetUser(id uint) (*model.User, error) {
	return &model.User{}, nil
}

func (r *repoInstance) UpdateUser(user *model.User) error {
	return nil
}

func (r *repoInstance) DeleteUser(id uint) error {
	return nil
}
