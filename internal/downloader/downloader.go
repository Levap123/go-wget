package download

import (
	"fmt"
	"go-wget/internal/service"
	utils "go-wget/pkg"
	"io"
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

func (d *Downloader) Download(url, path, filename string, w io.Writer) error {
	if filename == "" {
		filename = d.service.GetFilename(url)
	}
	
	fmt.Fprintf(w, "starts at %v\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(w, "sending request, awaiting response... ")

	resp, err := d.httpClient.Get(url)
	if resp == nil {
		return ErrResponceNil
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	filename = d.service.DefineFilename(filename, contentType)
	fmt.Fprintln(w, resp.Status)

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	if size != -1 {
		fmt.Fprintf(w, "content size: %d [~%fMB]\n", size, utils.ToMB(size))
	}

	fmt.Fprintf(w, "saving file to: %s\n", path+filename)

	// if size == -1 {
	// 	bar := pb.New64(0).SetRefreshRate(time.Second)
	// 	bar.Prefix(filename + "        ")
	// 	bar.Start()

	// 	reader := bar.NewProxyReader(resp.Body)
	// 	if err := d.service.GetFileWithContentLength(filename, path, reader); err != nil {
	// 		return err
	// 	}
	// 	bar.Finish()
	// 	return nil
	// }

	bar := pb.New(int(size)).SetRefreshRate(time.Second).SetUnits(pb.U_BYTES)
	bar.Output = w
	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	bar.Prefix(filename + "         ")

	bar.Start()
	reader := bar.NewProxyReader(resp.Body)
	if err := d.service.GetFileWithContentLength(filename, path, reader); err != nil {
		return err
	}
	bar.Finish()

	fmt.Fprintf(w, "\nDownloaded [%s]\n", url)
	fmt.Fprintf(w, "finished at %v\n", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
