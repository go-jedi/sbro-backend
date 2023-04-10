package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/sbro-backend/appl_row"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) SendMessageUser(userSendMessageForm appl_row.UserSendMessage) (int, error) {
	userSendMessageFormJson, err := json.Marshal(userSendMessageForm)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userSendMessageForm, %s", err)
	}
	_, err = r.db.Exec("SELECT user_send_message($1)", userSendMessageFormJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) SendDataUser(userSendDataForm appl_row.UserSendData) (int, error) {
	userSendDataFormJson, err := json.Marshal(userSendDataForm)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userSendDataForm, %s", err)
	}
	_, err = r.db.Exec("SELECT user_send_data($1)", userSendDataFormJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_send_data из базы данных, %s", err)
	}
	return http.StatusOK, nil
}
