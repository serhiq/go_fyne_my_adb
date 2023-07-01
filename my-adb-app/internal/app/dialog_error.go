package app

import "fyne.io/fyne/v2/dialog"

func (w *App) showError(e error) {
	dialog.ShowError(e, w.win)
}
