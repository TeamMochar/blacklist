package service

import "github.com/bostud/blacklist/pkg/repository"

type Authorithation interface {
}

type User interface {
}

type Complain interface {
}

type UserComplain interface {
}

type Service struct {
	Authorithation
	User
	Complain
	UserComplain
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
