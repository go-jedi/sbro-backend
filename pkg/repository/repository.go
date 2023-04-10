package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/sbro-backend/appl_row"
)

type TodoUser interface {
	SendMessageUser(userSendMessageForm appl_row.UserSendMessage) (int, error)
	SendDataUser(userSendDataForm appl_row.UserSendData) (int, error)
}

type Repository struct {
	TodoUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser: NewUserPostgres(db),
	}
}
