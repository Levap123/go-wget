package main

import (
	"flag"
	"fmt"
	"go-wget/internal/app"
	utils "go-wget/pkg"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("incorrect usage: go run . [OPTION]... [URL]...")
		return
	}

	var writer io.Writer
	path := flag.String("P", "", "path of the downloaded file")
	isLog := flag.Bool("B", false, "define output of downloaded file")
	filename := flag.String("O", "", "set filename of downloaded file")
	flag.Parse()

	if *isLog {
		file, err := os.Create("wget-log")
		if err != nil {
			log.Fatalln(err)
		}
		writer = file
		fmt.Println(`Output will be written to "wget-log".`)
		defer file.Close()
	} else {
		writer = os.Stdout
	}

	absPath, err := utils.ExpandPath(*path)
	if err != nil {
		log.Fatalln(err)
	}

	app, err := app.NewApp(500000000)
	if err != nil {
		log.Fatalln(err)
	}
	link := os.Args[len(os.Args)-1]

	if err := app.D.Download(link, absPath, *filename, writer); err != nil {
		log.Fatalln(err)
	}
}
