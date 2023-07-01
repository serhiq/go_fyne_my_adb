package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/serhiq/go_fyne_my_adb/internal/utils"
	"log"
)

func (s *App) showPickFolder() {
	fileDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			log.Println(err)
			return
		}
		if uri != nil {
			path := uri.Path()
			s.runtimeSettings.ScreenshotFolder = path
		}

	}, s.win)

	fileDialog.Resize(fyne.NewSize(1000, 500))

	if s.runtimeSettings.ScreenshotFolder != "" {
		path, err := utils.PathToUrl(s.runtimeSettings.ScreenshotFolder)
		if err != nil {
			log.Println(err)
		} else {
			fileDialog.SetLocation(path)
		}

	}

	fileDialog.Show()
}
