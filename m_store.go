package main

import (
	"fmt"
	"log"
	"os"
)

func getDirname() string {
	var dirname, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	var directory = fmt.Sprintf("%s/Pictures", dirname)
	log.Println("Directory", directory)
	return directory
}
