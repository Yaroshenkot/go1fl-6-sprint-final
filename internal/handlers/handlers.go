package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Metod not allowed", http.StatusMethodNotAllowed)
		return
	}
	content, err := os.ReadFile("../index.html")
	if err != nil {
		content = []byte("Read file error.") /*
					content = []byte(`<!DOCTYPE html>
			<html lang="en">
			  <head>
			    <meta charset="UTF-8" />
			    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
			    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
			    <title>Document</title>
			  </head>
			  <body>
			    <form
			      enctype="multipart/form-data"
			      action="http://localhost:8080/upload"
			      method="post"
			    >
			      <input type="file" name="myFile" />
			      <input type="submit" value="upload" />
			    </form>
			  </body>
			</html>`)*/
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Cannot parse form ", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Cannot get file from form ", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Cannot read file ", http.StatusInternalServerError)
		return
	}

	result, err := service.AutoDetectAndConvert(string(content))
	if err != nil {
		http.Error(w, "Conversion error ", http.StatusInternalServerError)
		return
	}

	filename := "result_" + time.Now().Format("20060102_150405") + filepath.Ext(header.Filename)
	if filename == "result_" {
		filename += ".txt"
	}
	os.WriteFile(filename, []byte(result), 0644)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))

}
