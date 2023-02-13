package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ToMB(value int) float64 {
	return float64(value) / (1024 * 1024)
}

func ExpandPath(dest string) (string, error) {
	if strings.HasPrefix(dest, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return strings.Replace(dest, "~", homeDir, 1) + "/", nil
	}

	abs, err := filepath.Abs(dest)
	if err != nil {
		return "", err
	}
	return abs + "/", nil
}
