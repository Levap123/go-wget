package main

import (
	"flag"
	"fmt"
	"go-wget/internal/app"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("incorrect usage: go run . [OPTION]... [URL]...")
	}

	path := flag.String("P", "", "path of the downloaded file")
	flag.Parse()

	absPath, err := expandPath(*path)
	if err != nil {
		log.Fatalln(err)
	}

	app := app.NewApp()
	link := os.Args[len(os.Args)-1]

	if err := app.D.Download(link, absPath); err != nil {
		log.Fatal(err)
	}
}

func expandPath(dest string) (string, error) {
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
