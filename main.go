package main

import (
	"flag"
	"fmt"
	"go-wget/internal/app"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("incorrect usage: go run . [OPTION]... [URL]...")
		return
	}

	var writer io.Writer
	path := flag.String("P", "", "path of the downloaded file")
	isLog := flag.Bool("B", false, "define output of downloaded file")
	flag.Parse()

	if *isLog {
		file, err := os.Create("wget-log")
		if err != nil {
			log.Fatalln(err)
		}
		writer = file
		fmt.Println(`Output will be written to "wget-log".`)
	} else {
		writer = os.Stdout
	}

	absPath, err := expandPath(*path)
	if err != nil {
		log.Fatalln(err)
	}

	app := app.NewApp()
	link := os.Args[len(os.Args)-1]

	if err := app.D.Download(link, absPath, writer); err != nil {
		log.Fatalln(err)
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
