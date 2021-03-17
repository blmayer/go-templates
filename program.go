package main

import (
	"os"

	"github.com/blmayer/template/internal/database/mongodb"
)

const help = `template is a go program template
Usage:
  template [options]
Available options:
  -h
  --help	show this help
Examples:
  template --help`

var nosql mongodb.Database

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

	// Connect to mongodb
	var err error
	nosql, err = mongodb.Connect("connString", "myDB")
	if err != nil {
		panic("mongodb connection: " + err.Error())
	}

	println("template finished")
}
