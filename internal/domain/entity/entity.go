package entity

import (
	"errors"
	"fmt"
	"strings"
)



type Link struct{
	ID int
	OriginalURL string
	ShortURL string
}


func NewLink(short_url,original_url string)(*Link, error){
	if checkURL(original_url) == false{
		return nil,errors.New(original_url)
	}
	return &Link{
		OriginalURL: normalizeURL(original_url),
		ShortURL: short_url,
	}, nil
}

func normalizeURL(original_link string)string{
	if strings.HasPrefix(original_link, "http://") || 
	strings.HasPrefix(original_link, "https://"){
		return original_link
	}
	return fmt.Sprintf("https://%s",original_link)
}

func checkURL(original_link string)bool{
	return original_link != ""
}