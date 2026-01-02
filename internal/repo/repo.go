package repo

import (
	"context"
	"errors"
	"log"
	"urlshortener/internal/domain/entity"
	errmsg "urlshortener/internal/errMsg"

	"gorm.io/gorm"
)


type Repository interface{
	Create(ctx context.Context,link *entity.Link)error
	Get(ctx context.Context)([]*entity.Link,error)
	GetByShortURL(ctx context.Context, shortURL string)(*entity.Link,error)
}

type Link struct{
	ID int `gorm:"primaryKey;autoIncrement"`
	Original_url string `gorm:"not null"`
	Short_url string `gorm:"unique;not null"`
}

type LinkRepository struct{
	Database *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository{
	return &LinkRepository{
		Database: db,
	}
}


func fromDomain(link *entity.Link)*Link{
	return &Link{
		Original_url: link.OriginalURL,
		Short_url: link.ShortURL,
	}
}

func toDomain(link *Link)*entity.Link{
	return &entity.Link{
		ID: link.ID,
		OriginalURL: link.Original_url,
		ShortURL: link.Short_url,
	}
}

func toDomains(links []*Link)[]*entity.Link{
	response:=make([]*entity.Link, len(links))
	for i, link := range links{
		response[i]=toDomain(link)
	}
	return response
}


func (r *LinkRepository) Create(ctx context.Context,link *entity.Link)error{
	err := r.Database.WithContext(ctx).Create(fromDomain(link)).Error
	if err != nil{
		log.Print(err)
		return errmsg.ErrFailedCreateLink
	}
	return nil
}

func (r *LinkRepository)Get(ctx context.Context)([]*entity.Link,error){
	var links []*Link
	err:= r.Database.WithContext(ctx).Find(&links).Error
	if err != nil{
		log.Print(err)
		return nil, errmsg.ErrFailedGetLink
	}
	return toDomains(links),nil
}

func (r *LinkRepository)GetByShortURL(ctx context.Context, shortURL string)(*entity.Link,error){ 
	var link Link
	err := r.Database.WithContext(ctx).Where("short_url = ?", shortURL).First(&link).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil,errmsg.ErrFailedGetLink
		}
		log.Print(err)
		return nil,errmsg.ErrFailedGetLink
	}
	return toDomain(&link),nil
}