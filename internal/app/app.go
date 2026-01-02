package app

import (
	"os"
	"urlshortener/internal/config"
	"urlshortener/internal/db"
	"urlshortener/internal/handlers"
	"urlshortener/internal/repo"
	"urlshortener/internal/service"
	"urlshortener/internal/tools"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func StartApp(){
	godotenv.Load()


	config := config.LoadConfig()
	db:=db.InitDB(config.DatabaseConfig.DSN())
	repo:=repo.NewLinkRepository(db.DB)
	generator := tools.NanoGenerator{}
	service:=service.NewLinkService(repo, &generator)
	handlers:=handlers.NewLinkHandlers(service)
	
	
	
	r:=gin.Default()
	r.Use(cors.Default())
	
	api:=r.Group("api")
	//Ручка создания ссылки
	api.POST("/create",func(c *gin.Context){
		handlers.CreateLink(c)
	})
	//Ручка получения всех ссылок 
	api.GET("/getallurl", func(c* gin.Context){
		handlers.GetLinks(c)
	})
	//Редирект на оригинальную ссылку
	r.GET("/:short",func(c *gin.Context){
		handlers.RedirectFromShortURL(c)
	})

	port:=os.Getenv("SERVICE_PORT")
	r.Run(":"+port)
}