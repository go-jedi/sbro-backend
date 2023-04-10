package service

import (
	"github.com/rob-bender/sbro-backend/appl_row"
	"github.com/rob-bender/sbro-backend/pkg/repository"
)

type TodoUser interface {
	SendMessageUser(userSendMessageForm appl_row.UserSendMessage) (int, error)
	SendDataUser(userSendDataForm appl_row.UserSendData) (int, error)
}

type Service struct {
	TodoUser
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoUser: NewUserService(r.TodoUser),
	}
}
