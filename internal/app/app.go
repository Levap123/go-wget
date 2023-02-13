package app

import (
	"errors"
	download "go-wget/internal/downloader"
	"go-wget/internal/service"
	"net/http"
	"regexp"
)

type App struct {
	D *download.Downloader
}

func NewApp(speedLimit int64) (*App, error) {
	if speedLimit < 0 {
		return nil, errors.New("speed limit cannot be zero ol less")
	}
	cl := &http.Client{Transport: &download.RateLimitedTransport{http.DefaultTransport, speedLimit}}
	re := regexp.MustCompile(`([^/]+)\.([^/?]+)(\?.*)?$`)
	service := service.NewService(re)
	downloader := download.NewDownloader(cl, service)
	return &App{
		D: downloader,
	}, nil
}
