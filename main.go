package main

import (
	"fmt"
	"go-wget/internal/app"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("incorrect usage: go run . [OPTION]... [URL]...")
	}
	app := app.NewApp()
	link := os.Args[1]
	if err := app.D.Download(link, ""); err != nil {
		log.Fatal(err)
	}
}
