package tools

import (
	"fmt"
	"log"
	"os"
	"strconv"
	errmsg "urlshortener/internal/errMsg"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Generator interface{
	GenerateUniqueID()(string,error)
}

type NanoGenerator struct{}

//Генерация уникального ID для БД
func(n *NanoGenerator) GenerateUniqueID()(string,error){
	alphabet := os.Getenv("ALPHABET_GEN")
	length := os.Getenv("LENGTH_GEN")
	host := os.Getenv("VIRTUAL_HOST")
	if alphabet == ""{
		alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	}
	if length == ""{
		length = "10"
	}
	l,err := strconv.Atoi(length)
	if err != nil{
		log.Print(err)
		return "",errmsg.ErrFailedConvertStr
	}
	shortUrl,err := gonanoid.Generate(alphabet,l)
	if err != nil{
		log.Print(err)
		return "",errmsg.ErrFailedCreateShort
	}
	return fmt.Sprintf("https://%s/%s",host,shortUrl),nil
}
