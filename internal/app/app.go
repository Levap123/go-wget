package app

import (
	download "go-wget/internal/downloader"
	"go-wget/internal/service"
	"net/http"
	"regexp"
)

type App struct {
	D *download.Downloader
}

func NewApp() *App {
	cl := &http.Client{}
	service := new(service.Service)
	re := regexp.MustCompile(`.*/(.*)`)
	downloader := download.NewDownloader(cl, service, re)
	return &App{
		D: downloader,
	}
}
