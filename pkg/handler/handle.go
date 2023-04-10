package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rob-bender/sbro-backend/pkg/service"

	cors "github.com/rs/cors/wrapper/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler { // создаём новый handler с полем services
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine { // обработчик роутов, Создание роутов
	router := gin.New() // инициализация роутов
	router.Use(cors.Default())

	api := router.Group("/api-v1")
	{
		api.POST("/user/sendMessage", h.sendMessageUser) // отправка сообщения
		api.POST("/user/sendData", h.sendDataUser)       // отправка данных
	}

	return router
}
