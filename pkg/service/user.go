package service

import (
	"github.com/rob-bender/sbro-backend/appl_row"
	"github.com/rob-bender/sbro-backend/pkg/repository"
)

type UserService struct {
	repo repository.TodoUser
}

func NewUserService(r repository.TodoUser) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) SendMessageUser(userSendMessageForm appl_row.UserSendMessage) (int, error) {
	return s.repo.SendMessageUser(userSendMessageForm)
}

func (s *UserService) SendDataUser(userSendDataForm appl_row.UserSendData) (int, error) {
	return s.repo.SendDataUser(userSendDataForm)
}
