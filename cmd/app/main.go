package main

import (
	"urlshortener/internal/app"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	app.StartApp()
}