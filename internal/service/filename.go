package service

import "strings"

func (s *Service) DefineFilename(filename, contentType string) string {
	if filename == "" {
		switch {
		case strings.Contains(contentType, "html"):
			return "index.html"

		case strings.Contains(contentType, "x-icon"):
			return "favicon.ico"

		case strings.Contains(contentType, "png"):
			return "image.png"

		case strings.Contains(contentType, "jpeg"):
			return "image.jpeg"

		case strings.Contains(contentType, "pdf"):
			return "file.pdf"

		case strings.Contains(contentType, "plain"):
			return "file.txt"

		}

	}
	return filename
}
