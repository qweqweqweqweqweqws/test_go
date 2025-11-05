# REST API Backend на Go

Проект представляет собой backend REST API, построенный на Go с использованием фреймворка Gin.

## Структура проекта

```
back/
├── internal/
│   ├── config/      # Конфигурация приложения
│   ├── handlers/    # HTTP обработчики
│   ├── middleware/  # Middleware (CORS, Logger, Recovery)
│   └── routes/      # Определение маршрутов
├── main.go          # Точка входа приложения
├── go.mod           # Зависимости проекта
└── README.md        # Документация
```

## Требования

- Go 1.21 или выше
- Git

## Установка и запуск

1. Клонируйте репозиторий (если применимо):
```bash
git clone <repository-url>
cd back
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл `.env` на основе `.env.example`:
```bash
cp .env.example .env
```

4. Запустите сервер:
```bash
go run main.go
```

Сервер будет доступен по адресу: `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /api/v1/health` - Проверка статуса API

### Примеры
- `GET /api/v1/example` - Пример GET запроса
- `POST /api/v1/example` - Пример POST запроса

### Корневой маршрут
- `GET /` - Информация об API

## Примеры запросов

### Health Check
```bash
curl http://localhost:8080/api/v1/health
```

### GET Example
```bash
curl http://localhost:8080/api/v1/example
```

### POST Example
```bash
curl -X POST http://localhost:8080/api/v1/example \
  -H "Content-Type: application/json" \
  -d '{"name":"Иван","email":"ivan@example.com"}'
```

## Переменные окружения

- `PORT` - Порт для запуска сервера (по умолчанию: 8080)
- `ENVIRONMENT` - Окружение (development/production)
- `DATABASE_URL` - URL базы данных (если используется)

## Разработка

Для сборки проекта:
```bash
go build -o bin/api main.go
```

Для запуска тестов:
```bash
go test ./...
```

## Используемые библиотеки

- [Gin](https://github.com/gin-gonic/gin) - HTTP веб-фреймворк
- [godotenv](https://github.com/joho/godotenv) - Загрузка переменных окружения из .env файла

## Лицензия

MIT

