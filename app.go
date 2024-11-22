package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
// func (a *App) startup(ctx context.Context) {
// 	a.ctx = ctx
// }

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SetFolderPath() string {
	folderPath, _ := runtime.OpenDirectoryDialog(custom_context, runtime.OpenDialogOptions{
		DefaultDirectory: getHomeDirPath(),
		ShowHiddenFiles:  false,
	})

	return folderPath
}

func (a *App) GetImages(path string) Images {
	return updateJsonImages(path)
}

func (a *App) OpenInFileExplorer(path string) error {
	err := _openInFileExplorer(path)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) AddTag(location string, imageName string, addedTag string) Images {
	return _addTag(location, imageName, addedTag)
}

func (a *App) RemoveTag(location string, imageName string, removedTag string) Images {
	return _removeTag(location, imageName, removedTag)
}

func (a *App) AddImageByLink(location string, link string, newTag string) Images {
	return _addImageByLink(location, link, newTag)
}

func (a *App) DeleteImageWithLink(location string, name string) Images {
	return _deleteImageWithLink(location, name)
}
