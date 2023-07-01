package app

import (
	"fyne.io/fyne/v2/widget"
	"github.com/serhiq/go_fyne_my_adb/internal/commands"
	"time"
)

func (w *App) AddCmd(cmd *commands.Cmd) {
	w.cmds = append(w.cmds, cmd)
	w.cmdbar.Add(widget.NewButton(cmd.Name, func() {
		w.sendCmd(cmd)
	}))
}

func (w *App) sendCmd(cmd *commands.Cmd) {

	if cmd.UseSelectedList {

		for name, _ := range w.packageNameList {
			_, err := w.MyTerm.Term.Write([]byte(cmd.Command + " " + name + "\n"))
			if err != nil {
				println("err")

			}
			time.Sleep(1000)
		}

		return

	}

	_, err := w.MyTerm.Term.Write([]byte(cmd.Command + "\n"))
	if err != nil {
		println("err")

	}
}

func (w *App) addPackageToList(packageApp string, include bool) {
	_, has := w.packageNameList[packageApp]
	if include && has {
		return
	}

	if include && !has {
		w.packageNameList[packageApp] = struct{}{}
	}

	if !include && has {
		delete(w.packageNameList, packageApp)
	}
}
