package handler

import (
	"github.com/VeeRomanoff/todo-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// инициализация всех эндпоинтов
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllLists)
			list.GET("/:uuid", h.getListByListUUID)
			list.PUT("/:uuid", h.updateListByListUUID)
			list.DELETE("/:uuid", h.deleteListByListUUID)
		}

		items := api.Group(":uuid/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item_uuid", h.getItemByItemUUID)
			items.PUT("/:item_uuid", h.updateItemByItemUUID)
			items.DELETE("/:item_uuid", h.deleteItemByItemUUID)
		}
	}

	return router
}
