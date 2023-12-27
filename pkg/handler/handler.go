package handler

import (
	"github.com/bostud/blacklist/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		complains := api.Group("/complains")
		{
			complains.POST("/", h.createComplain)
			complains.GET("/:id", h.getComplain)
			complains.PUT("/:id", h.updateComplain)
			complains.DELETE("/:id", h.deleteComplain)
			complains.GET("/search", h.searchComplains)
		}
		users := api.Group("/users")
		{
			users.GET("/:id", h.getUser)
		}

		userComplains := api.Group("/users-complains")
		{
			userComplains.GET("/:userId", h.getUsersComplains)
		}
	}

	return router
}
