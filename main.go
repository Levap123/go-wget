package main

import (
	"fmt"
	"go-wget/internal/app"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("incorrect usage: go run . [OPTION]... [URL]...")
	}
	app := app.NewApp()
	link := os.Args[1]
	app.D.Download(link, "")
}
