package main

import (
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

	println("template finished")
}
