package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
)

func getHomeDirPath() string {
	var dirname, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}

func _openInFileExplorer(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	err := clipboard.WriteAll(path)
	if err != nil {
		return err
	}

	/* TODO: add for windows */
	if _, err := exec.LookPath("xdg-open"); err != nil {
		return err
	}

	cmd := exec.Command("xdg-open", path)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
