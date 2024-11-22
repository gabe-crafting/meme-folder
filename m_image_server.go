package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func ServeImages(w http.ResponseWriter, r *http.Request, imagesDir string) {
	imageName := filepath.Base(r.URL.Path)
	imagePath := filepath.Join(imagesDir, imageName)

	fmt.Println(imagePath)
}
