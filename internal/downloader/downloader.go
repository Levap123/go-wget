package download

import (
	"go-wget/internal/service"
	"net/http"
	"regexp"
)

type Downloader struct {
	service        *service.Service
	httpClient     *http.Client
	filenameRegexp *regexp.Regexp
}

func NewDownloader(cl *http.Client, service *service.Service, re *regexp.Regexp) *Downloader {
	return &Downloader{
		service:        service,
		httpClient:     cl,
		filenameRegexp: re,
	}
}

func (d *Downloader) Download(url, path string) error {
	filename := d.filenameRegexp.FindStringSubmatch(url)
	resp, err := d.httpClient.Get(url)
	if err != nil {
		return err
	}
	if err := d.service.GetFile(filename[1], path, resp.Body); err != nil {
		return err
	}
	return nil
}
