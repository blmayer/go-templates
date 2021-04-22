package main

import (
	"mime"
	"net/http"
	"strings"
	"time"
)

type templateData struct {
	DayTime string
	Color   string
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "" {
		path = "index.html"
	}

	example := templateData{
		DayTime: time.Now().Format(time.RFC1123),
		Color:   "blue",
	}

	ext := mime.TypeByExtension(path[strings.LastIndex(path, "."):])
	w.Header()["Content-Type"] = []string{ext}

	err := temp.ExecuteTemplate(w, path, example)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web"+r.URL.Path)
}
