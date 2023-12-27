package repository

type Authorithation interface {
}

type User interface {
}

type Complain interface {
}

type UserComplain interface {
}

type Repository struct {
	Authorithation
	User
	Complain
	UserComplain
}

func Newrepository() *Repository {
	return &Repository{}
}
