package handlers

import (
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "couldn't open the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}
func UploaderHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "couldn't parse the form", http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "couldn't read the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "couldn't read the file", http.StatusInternalServerError)
		return
	}
	converted, err := service.Converter(string(data))
	if err != nil {
		http.Error(w, "couldn't convert the file", http.StatusInternalServerError)
	}
	ext := filepath.Ext(header.Filename)
	t := time.Now().UTC().Format("01022006_150405")
	fileName := fmt.Sprintf("%s%s", t, ext)
	if err := os.WriteFile(fileName, []byte(converted), 0644); err != nil {
		http.Error(w, "couldn't write the file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
