package download

import (
	"go-wget/internal/service"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/cheggaaa/pb"
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
	if resp == nil {
		return ErrResponceNil
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	bar := pb.New(int(size)).SetRefreshRate(time.Second * 5)

	bar.Start()
	reader := bar.NewProxyReader(resp.Body)
	if err := d.service.GetFileWithContentLength(filename[1], path, reader); err != nil {
		return err
	}
	bar.Finish()

	return nil
}
