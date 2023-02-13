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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(path + name)
	if err != nil {
		return fmt.Errorf("fail in creating file - %w", err)
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

func (s *Service) GetFileWithoutContentLength(name, path string, src io.Reader) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(path + name)
	if err != nil {
		return fmt.Errorf("fail in creating file - %w", err)
	}

	if _, err := io.Copy(file, src); err != nil {
		return fmt.Errorf("fail in copy file - %w", err)
	}

	defer func() {
		file.Close()
	}()
	return nil
}
