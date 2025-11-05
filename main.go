package main

import (
	"log"
	"os"

	"back/internal/config"
	"back/internal/handlers"
	"back/internal/middleware"
	"back/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, используем переменные окружения системы")
	}

	// Загружаем конфигурацию
	cfg := config.Load()

	// Устанавливаем режим Gin
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Создаем роутер
	router := gin.Default()

	// Подключаем middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	// Инициализируем обработчики
	h := handlers.New()

	// Настраиваем маршруты
	routes.SetupRoutes(router, h)

	// Получаем порт из переменной окружения или используем дефолтный
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Сервер запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
