package service

import (
	"fmt"
	"io"
	"os"
)

type Service struct{}

func (s *Service) GetFile(name, path string, src io.Reader) error {
	file, err := os.Create(path + name)
	if err != nil {
		return fmt.Errorf("fail in getting file - %w", err)
	}

	if _, err := io.Copy(file, src); err != nil {
		return fmt.Errorf("fail in copy file - %w", err)
	}
	return nil
}