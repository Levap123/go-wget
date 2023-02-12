package download

import (
	"fmt"
	"go-wget/internal/service"
	utils "go-wget/pkg"
	"net/http"
	"strconv"
	"time"

	"github.com/cheggaaa/pb"
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

func (d *Downloader) Download(url, path string) error {
	filename := d.service.GetFilename(url)

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

	if filename == "" {
		// // TODO
		// // GET FILENAME FROM CONTENT TYPE
		// contentType := resp.Header.Get("Content-Type")
		// ext, err := mime.ExtensionsByType(contentType)
		// if err != nil {
		// 	return err
		// }
		// fmt.Println(ext)
		filename = "test"
	}

	fmt.Println(resp.Status)

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	if size != -1 {
		fmt.Printf("content size: %d [~%fMB]\n", size, utils.ToMB(size))
	}

	fmt.Printf("saving file to: %s\n", path+filename)
	var bar *pb.ProgressBar
	if size > 0 {
		bar = pb.New(size).SetRefreshRate(time.Second).SetUnits(pb.U_BYTES)
	} else {
		bar = pb.New(0).SetRefreshRate(time.Second).SetUnits(pb.U_BYTES)
	}

	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	bar.Prefix(filename + "         ")

	bar.Start()
	reader := bar.NewProxyReader(resp.Body)
	if err := d.service.GetFileWithContentLength(filename, path, reader); err != nil {
		return err
	}
	bar.Finish()

	fmt.Printf("\nDownloaded [%s]\n", url)
	fmt.Printf("finished at %v\n", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
