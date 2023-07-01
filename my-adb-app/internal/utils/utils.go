package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"path/filepath"
)

func PathToUrl(path string) (fyne.ListableURI, error) {
	path = filepath.ToSlash(path)

	uri, err := storage.ParseURI("file://" + path)
	if err != nil {
		return nil, err
	}

	return storage.ListerForURI(uri)
}
