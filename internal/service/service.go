package service

import (
	"context"
	"urlshortener/internal/domain/entity"
	"urlshortener/internal/dto"
	"urlshortener/internal/repo"
	"urlshortener/internal/tools"
)

type Service interface{
	NewLink(ctx context.Context,dto *dto.NewLink)(string,string,error)
	GetLinks(ctx context.Context)([]*entity.Link,error)
	GetLink(ctx context.Context,shortURL string )(string,error)
}


type LinkService struct{
	repo repo.Repository
	generator tools.Generator
}

func NewLinkService(repo repo.Repository, generator tools.Generator)*LinkService{
	return &LinkService{
		repo: repo,
		generator: generator,
	}
}

func (s *LinkService) NewLink(ctx context.Context,dto *dto.NewLink)(string,string,error){
	shortUrl,err := s.generator.GenerateUniqueID()
	if err != nil{
		return "","",err
	}
	newLink,err := entity.NewLink(shortUrl,dto.OriginalURL)
	if err != nil{
		return "","",err
	}
	err = s.repo.Create(ctx,newLink)
	if err != nil{
		return "","",err
	}
	return dto.OriginalURL, shortUrl, nil
}

func (s *LinkService) GetLinks(ctx context.Context)([]*entity.Link,error){
	links,err := s.repo.Get(ctx)
	if err!=nil{
		return nil,err
	}
	return links,nil
}

func (s *LinkService) GetLink(ctx context.Context,shortURL string  )(string,error){
	link,err := s.repo.GetByShortURL(ctx,shortURL)
	if err!=nil{
		return "",err
	}
	return link.OriginalURL,nil
}


