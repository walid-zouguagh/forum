package controllers

import (
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func Server(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "..") {
		http.Error(w, "Page not Found", http.StatusNotFound)
		return
	}

	filename := "." + path.Clean(r.URL.Path)
	file, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Page not Found", http.StatusNotFound)
		return
	}
	http.ServeContent(w, r, filename, time.Now(), strings.NewReader(string(file)))
}
