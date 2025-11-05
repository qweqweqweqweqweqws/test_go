package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	// Здесь можно добавить зависимости (например, сервисы, репозитории)
}

func New() *Handlers {
	return &Handlers{}
}

// HealthCheck проверяет статус API
func (h *Handlers) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "API работает",
	})
}

// GetExample пример обработчика GET запроса
func (h *Handlers) GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Пример GET запроса",
		"data":    []string{"элемент 1", "элемент 2", "элемент 3"},
	})
}

// PostExample пример обработчика POST запроса
func (h *Handlers) PostExample(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Неверный формат данных",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Данные успешно созданы",
		"data":    req,
	})
}

// cookieParams структура для параметров cookie
type cookieParams struct {
	Name     string
	Value    string
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
	SameSite string
}

// setCookieHelper устанавливает cookie с заданными параметрами
func (h *Handlers) setCookieHelper(c *gin.Context, params cookieParams) {
	// Устанавливаем значения по умолчанию
	if params.Path == "" {
		params.Path = "/"
	}
	if params.MaxAge == 0 {
		params.MaxAge = 3600 // 1 час по умолчанию
	}

	// Проверяем, является ли запрос cross-origin
	origin := c.Request.Header.Get("Origin")
	isCrossOrigin := origin != ""

	// Для cross-origin запросов используем SameSite=None и Secure=true
	if params.SameSite == "" {
		if isCrossOrigin {
			params.SameSite = "None"
			params.Secure = true // Обязательно для SameSite=None
		} else {
			params.SameSite = "Lax"
		}
	} else if params.SameSite == "None" && !params.Secure {
		// Если SameSite=None, Secure должен быть true
		params.Secure = true
	}

	// Определяем SameSite
	var sameSite http.SameSite
	switch params.SameSite {
	case "Strict":
		sameSite = http.SameSiteStrictMode
	case "Lax":
		sameSite = http.SameSiteLaxMode
	case "None":
		sameSite = http.SameSiteNoneMode
		// Для SameSite=None Secure должен быть true
		if !params.Secure {
			params.Secure = true
		}
	default:
		sameSite = http.SameSiteLaxMode
	}

	// Устанавливаем cookie с поддержкой SameSite
	cookie := &http.Cookie{
		Name:     params.Name,
		Value:    params.Value,
		Path:     params.Path,
		Domain:   params.Domain,
		MaxAge:   params.MaxAge,
		Secure:   params.Secure,
		HttpOnly: params.HttpOnly,
		SameSite: sameSite,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cookie успешно установлен",
		"cookie": gin.H{
			"name":      params.Name,
			"value":     params.Value,
			"max_age":   params.MaxAge,
			"path":      params.Path,
			"domain":    params.Domain,
			"secure":    params.Secure,
			"http_only": params.HttpOnly,
			"same_site": params.SameSite,
		},
	})
}

// SetCookie устанавливает cookie через POST запрос (JSON body)
func (h *Handlers) SetCookie(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Value    string `json:"value" binding:"required"`
		MaxAge   int    `json:"max_age"`   // В секундах (0 = сессионная cookie)
		Path     string `json:"path"`      // Путь (по умолчанию "/")
		Domain   string `json:"domain"`    // Домен
		Secure   bool   `json:"secure"`    // HTTPS only
		HttpOnly bool   `json:"http_only"` // Доступ только через HTTP (не JS)
		SameSite string `json:"same_site"` // Strict, Lax, None
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Неверный формат данных",
			"details": err.Error(),
		})
		return
	}

	h.setCookieHelper(c, cookieParams{
		Name:     req.Name,
		Value:    req.Value,
		MaxAge:   req.MaxAge,
		Path:     req.Path,
		Domain:   req.Domain,
		Secure:   req.Secure,
		HttpOnly: req.HttpOnly,
		SameSite: req.SameSite,
	})
}

// SetCookieGET устанавливает cookie через GET запрос (query parameters)
func (h *Handlers) SetCookieGET(c *gin.Context) {
	name := c.Query("name")
	value := c.Query("value")

	if name == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Параметры 'name' и 'value' обязательны",
		})
		return
	}

	var maxAge int
	if m := c.Query("max_age"); m != "" {
		if parsed, err := strconv.Atoi(m); err == nil {
			maxAge = parsed
		}
	}

	var secure bool
	if s := c.Query("secure"); s == "true" || s == "1" {
		secure = true
	}

	var httpOnly bool
	if h := c.Query("http_only"); h == "true" || h == "1" {
		httpOnly = true
	}

	h.setCookieHelper(c, cookieParams{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Path:     c.DefaultQuery("path", "/"),
		Domain:   c.Query("domain"),
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: c.DefaultQuery("same_site", ""),
	})
}
