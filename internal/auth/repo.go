package auth

import "go-web-staging/internal/entity"

type Repo interface {
	Create(user *entity.User) error
}

type repo struct{}

func NewRepo() Repo {
	return &repo{}
}

func (r *repo) Create(user *entity.User) error {
	return nil
}
