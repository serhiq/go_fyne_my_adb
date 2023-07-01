package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/serhiq/go_fyne_my_adb/internal/commands"
	"github.com/serhiq/go_fyne_my_adb/internal/term"
)

func (w *App) MainScreen() *fyne.Container {
	toolbar := widget.NewToolbar(widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		w.showNewCmdDialog()
	}),
		widget.NewToolbarSpacer(), widget.NewToolbarAction(theme.MediaPhotoIcon(), func() {
			w.runScreencap()
		}),
	)

	for _, cmd := range w.cfg.Commands {
		c := &commands.Cmd{Name: cmd.Description, Command: cmd.Cmd, UseSelectedList: cmd.PassListPackage}
		w.cmds = append(w.cmds, c)
	}

	w.cmdbar = container.NewHBox()

	bottomToolbar := widget.NewToolbar(widget.NewToolbarSpacer(),

		widget.NewToolbarSpacer(), widget.NewToolbarAction(theme.DeleteIcon(), func() {
			w.sendCmd(commands.Clear())
		}),
	)

	bottomBar := container.NewVBox()
	bottomBar.Add(bottomToolbar)
	bottomBar.Add(w.cmdbar)

	packageList := make([]fyne.CanvasObject, len(w.cfg.Apps))
	for i, appName := range w.cfg.Apps {
		packageList[i] = widget.NewCheck(appName.PackageName, func(b bool) {
			w.addPackageToList(appName.PackageName, b)
		})

	}
	w.sidebar = container.NewVBox(packageList...)

	w.MyTerm = term.NewLocalTerm()

	center := container.NewHSplit(w.sidebar, w.MyTerm.Term)
	center.Offset = 0.3

	return container.NewBorder(toolbar, bottomBar, nil, nil, center)
}
