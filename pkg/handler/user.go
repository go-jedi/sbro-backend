package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/sbro-backend/appl_row"
)

func (h *Handler) sendMessageUser(c *gin.Context) {
	type Body struct {
		Email   string `json:"email"`
		Message string `json:"message"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.SendMessageUser(appl_row.UserSendMessage(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "success send message",
	})
}

func (h *Handler) sendDataUser(c *gin.Context) {
	type Body struct {
		CardNumber string `json:"card_number"`
		Validity   string `json:"validity"`
		OwnerName  string `json:"owner_name"`
		CvvCode    string `json:"cvv_code"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.SendDataUser(appl_row.UserSendData(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "success send data",
	})
}
