package service

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

type Service struct {
	filenameRegexp *regexp.Regexp
}

func NewService(re *regexp.Regexp) *Service {
	return &Service{
		filenameRegexp: re,
	}
}

func (s *Service) GetFilename(url string) string {
	match := s.filenameRegexp.FindStringSubmatch(url)
	if len(match) >= 3 {
		return match[0]
	}
	return ""
}

func (s *Service) GetFileWithContentLength(name, path string, src io.ReadCloser) error {
	file, err := os.Create(path + name)
	if err != nil {
		return fmt.Errorf("fail in getting file - %w", err)
	}

	if _, err := io.Copy(file, src); err != nil {
		return fmt.Errorf("fail in copy file - %w", err)
	}

	defer func() {
		file.Close()
		src.Close()
	}()
	return nil
}
