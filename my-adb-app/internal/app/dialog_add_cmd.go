package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/serhiq/go_fyne_my_adb/internal/commands"
)

func (w *App) showNewCmdDialog() {
	nameEntry := widget.NewEntry()
	textEntry := widget.NewEntry()
	textEntry.MultiLine = true

	dlg := dialog.NewForm("New Command", "OK", "Cancel", []*widget.FormItem{
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Command", textEntry),
	}, func(b bool) {
		if b {
			cmd := &commands.Cmd{Name: nameEntry.Text, Command: textEntry.Text, Icon: ""}
			w.AddCmd(cmd)
		}
	}, w.win)

	dlg.Resize(fyne.NewSize(400, 300))
	dlg.Show()
}
