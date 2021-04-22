package main

import (
	"embed"
	"html/template"
	"net/http"
	"os"
)

const help = `template is a go program template
Usage:
  template [options]
Available options:
  -h
  --help	show this help
Examples:
  template --help`

var (
	temp *template.Template

	//go:embed web/*
	pages embed.FS
)

func main() {
	// Command line arguments parsing
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h":
			fallthrough
		case "--help":
			println(help)
			os.Exit(0)
		default:
			os.Stderr.WriteString("error: wrong argument\n")
			println(help)
			os.Exit(-1)
		}
	}

	var err error
	temp, err = template.ParseFS(pages, "web/*.*")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	http.HandleFunc("/images/", imagesHandler)

	http.ListenAndServe(":"+port, nil)
}
