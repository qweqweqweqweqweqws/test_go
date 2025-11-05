package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery обрабатывает паники и возвращает JSON ответ
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Внутренняя ошибка сервера",
			"message": "Произошла непредвиденная ошибка",
		})
		c.Abort()
	})
}
