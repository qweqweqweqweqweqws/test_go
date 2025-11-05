package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS настраивает заголовки CORS для работы с cross-origin запросами
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Разрешаем конкретный origin из запроса
		// Если origin не указан (например, same-origin запрос), разрешаем все
		// if origin != "" {
		// 	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// 	// Важно: если используем конкретный origin, можно использовать credentials
		// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// } else {
		// Если Origin не указан, разрешаем все (для same-origin запросов)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// Нельзя использовать credentials с "*"
		// }

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
