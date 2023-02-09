package download

import (
	"fmt"
	"go-wget/internal/service"
	utils "go-wget/pkg"
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

	fmt.Printf("starts at %v\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("sending request, awaiting response... ")

	resp, err := d.httpClient.Get(url)
	if resp == nil {
		return ErrResponceNil
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	fmt.Printf("content size: %d [~%fMB]\n", size, utils.ToMB(size))
	fmt.Printf("saving file to: %s\n", path+filename[1])

	bar := pb.New(int(size)).SetRefreshRate(time.Second).SetUnits(pb.U_BYTES)
	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	bar.Prefix(filename[1] + "         ")

	bar.Start()
	reader := bar.NewProxyReader(resp.Body)
	if err := d.service.GetFileWithContentLength(filename[1], path, reader); err != nil {
		return err
	}
	bar.Finish()

	fmt.Printf("\nDownloaded [%s]\n", url)
	fmt.Printf("finished at %v\n", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
