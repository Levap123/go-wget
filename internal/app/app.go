package app

import (
	download "go-wget/internal/downloader"
	"net/http"
)

type App struct {
	D *download.Downloader
}

func NewApp() *App {
	cl := &http.Client{}
	downloader := download.NewDownloader(cl)
	return &App{
		D: downloader,
	}
}
