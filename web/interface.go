package web

import entity "Concurrencia/internal/app/user/entities"

type UserService interface {
	CreateUser(user *entity.User) error
	GetUser(id string) (*entity.User, error)
}

type MsgService interface {
	PublishTopic(msg *entity.Msg)
}
