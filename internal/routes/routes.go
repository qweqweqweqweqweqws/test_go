package routes

import (
	"back/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	// API v1 группа
	v1 := router.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", h.HealthCheck)

		// Примеры эндпоинтов
		v1.GET("/example", h.GetExample)
		v1.POST("/example", h.PostExample)

		// Cookie эндпоинты
		v1.GET("/cookie/set", h.SetCookieGET) // GET с query параметрами
		v1.POST("/cookie/set", h.SetCookie)   // POST с JSON body

		// Здесь можно добавить другие маршруты
		// v1.GET("/users", h.GetUsers)
		// v1.POST("/users", h.CreateUser)
		// v1.GET("/users/:id", h.GetUserByID)
		// v1.PUT("/users/:id", h.UpdateUser)
		// v1.DELETE("/users/:id", h.DeleteUser)
	}

	// Корневой маршрут
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Добро пожаловать в REST API",
			"version": "1.0.0",
		})
	})
}
