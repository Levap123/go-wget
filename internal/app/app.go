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
	re := regexp.MustCompile(`([^/]+)\.([^/?]+)(\?.*)?$`)
	service := service.NewService(re)
	downloader := download.NewDownloader(cl, service)
	return &App{
		D: downloader,
	}
}
