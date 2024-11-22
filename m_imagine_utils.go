package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var imagePattern = regexp.MustCompile(`(?i)\.(jpg|jpeg|png|gif|bmp|tiff)$`)

// isImage checks if the given name is an image file and not a directory
func isImage(file os.DirEntry) bool {
	// Check if the name matches the image pattern and is not a directory
	return imagePattern.MatchString(file.Name()) && !file.IsDir()
}

type Image struct {
	Id   string   `json:"Id"`
	Name string   `json:"Name"`
	Path string   `json:"Path"`
	Tags []string `json:"Tags"`
	Link string   `json:"Link"`
}

type Images map[string]Image

func getFolderImages(dirTarget string) Images {
	images := Images{}

	files, err := os.ReadDir(dirTarget)

	if err != nil {
		log.Println("Inexistent file, opening default", dirTarget)
		log.Println(err)
		return Images{}
	}

	for _, file := range files {
		if isImage(file) {
			imagePath := fmt.Sprintf("%s/%s", dirTarget, file.Name())
			image := Image{
				Id:   file.Name(),
				Name: file.Name(),
				Path: imagePath,
				Tags: []string{},
			}
			images[file.Name()] = image
		}
	}

	return images
}

func writeJsonImages(data Images, location string) {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Error marshaling the data", data)
	}

	file, err := os.Create(fmt.Sprintf("%s/meme-folder.json", location))
	if err != nil {
		log.Println("Error create meme-folder.json:", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Println("Error write meme-folder.json:", err)
	}

	defer file.Close()
}

func readJsonImages(filePath string) Images {
	filePath = fmt.Sprintf("%s/meme-folder.json", filePath)

	byteValue, err := os.ReadFile(filePath)

	if err != nil {
		log.Println("reading-file:meme-folder.json", err)
		return Images{}
	}

	var images Images
	json.Unmarshal(byteValue, &images)

	if len(images) == 0 {
		return Images{}
	}

	return images
}

func updateJsonImages(location string) Images {

	if location == "" {
		log.Println("filepath empty: going to default")
		location = getDirname()
	}

	jsonImages := readJsonImages(location)
	folderImages := getFolderImages(location)

	for key, image := range folderImages {
		if _, exist := jsonImages[key]; !exist {
			log.Println("key:", key, "does exist")
			jsonImages[key] = image
		}
	}

	for key, image := range jsonImages {
		if _, exist := folderImages[key]; !exist && image.Link == "" {
			delete(jsonImages, key)
		}
	}

	writeJsonImages(jsonImages, location)
	return jsonImages
}

func _addTag(location string, imageName string, addedTag string) Images {

	if location == "" {
		log.Println("filepath empty: going to default")
		location = getDirname()
	}

	images := readJsonImages(location)
	if image, exists := images[imageName]; exists {
		newTags := addIfNotExists(image.Tags, addedTag)
		image.Tags = newTags
		images[imageName] = image
	}
	writeJsonImages(images, location)
	return images
}

func _removeTag(location string, imageName string, removedTag string) Images {
	if location == "" {
		location = getDirname()
	}

	images := readJsonImages(location)
	if image, exists := images[imageName]; exists {
		newTags := removeString(image.Tags, removedTag)
		image.Tags = newTags
		images[imageName] = image
	}
	writeJsonImages(images, location)
	return images
}

func _addImageByLink(location string, link string, newTag string) Images {
	if location == "" {
		location = getDirname()
	}

	var imageName string
	lastSlashIndex := strings.LastIndex(link, "/")
	if lastSlashIndex != -1 {
		imageName = link[lastSlashIndex+1:]
	}

	var newTags []string
	if newTag != "" {
		newTags = []string{"linked", newTag}
	} else {
		newTags = []string{"linked"}
	}

	images := readJsonImages(location)
	newImage := Image{
		Name: imageName,
		Path: link,
		Link: link,
		Id:   uuid.New().String(),
		Tags: newTags,
	}

	images[newImage.Id] = newImage
	writeJsonImages(images, location)
	return images
}

func _deleteImageWithLink(location string, name string) Images {
	if location == "" {
		log.Println("filepath empty: going to default")
		location = getDirname()
	}

	images := readJsonImages(location)
	delete(images, name)

	writeJsonImages(images, location)
	return images
}
