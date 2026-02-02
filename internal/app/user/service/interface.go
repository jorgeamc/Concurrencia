package service

import entity "Concurrencia/internal/app/user/entities"

type UserRepository interface {
	FindUserById(id string) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type RepositoryOne interface {
	GetRepositoryOne() string
}

type RepositoryTwo interface {
	GetRepositoryTwo() string
}

type MsgRepository interface {
	Publish(msg *entity.Msg)
}
