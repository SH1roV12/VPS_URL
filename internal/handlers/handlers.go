package handlers

import (
	"net/http"
	"urlshortener/internal/dto"
	"urlshortener/internal/service"

	"github.com/gin-gonic/gin"
)




type LinkHandlers struct{
	service service.Service
}

func NewLinkHandlers(service service.Service)*LinkHandlers{
	return &LinkHandlers{service: service}
}

//Создание короткой ссылки из большой входящей
func(h *LinkHandlers) CreateLink(c *gin.Context){
	var req dto.NewLink
	//Обработка входящей ссылки в формате JSON
	if err := c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	originalURL, shortURL, err := h.service.NewLink(c.Request.Context(),&req)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Ответ пользователю о добавлении ссылки
	c.JSON(http.StatusOK, gin.H{"original_url":originalURL,"short_url":shortURL})
}


//Получение всех данных о сохраненных ссылках 
func(h *LinkHandlers) GetLinks(c *gin.Context){
	links,err:=h.service.GetLinks(c.Request.Context())
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,links)
}


func(h *LinkHandlers) RedirectFromShortURL(c *gin.Context){
	//Поиск по какому ID искать оригинальную ссылку в БД
	shortUrl := c.Param("short")
	var originalURL string
	//Поиск оригинальной ссылки по короткой ссылке
	originalURL, err := h.service.GetLink(c.Request.Context(), shortUrl)
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
        return
    }
	//Редирект на оригинальную ссылку
	c.Redirect(http.StatusFound,originalURL)
}