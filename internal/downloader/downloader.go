package download

import (
	"go-wget/internal/service"
	"net/http"
)

type Downloader struct {
	service    *service.Service
	httpClient *http.Client
}

func NewDownloader(cl *http.Client, service *service.Service) *Downloader {
	return &Downloader{
		service:    service,
		httpClient: cl,
	}
}

func (d *Downloader) Download(url string)
