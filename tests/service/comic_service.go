package service

import (
	"node_puppeteer_example_go/comic_service"
)

type Comic interface {
	GetList() ([]*models.Comics, error)
	Count(int, error)
	GetInfo() (*models.Comics, error)
}

func NewComicService() *Comic {
	comicService := comic_service.ComicParam{
		PageNum:  1,
		PageSize: 10,
		Channel:  2,
		SourceID: 5830,
	}
	return &comicService
}
