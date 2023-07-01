package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/serhiq/go_fyne_my_adb/internal/commands"
	"github.com/serhiq/go_fyne_my_adb/internal/config"
	"github.com/serhiq/go_fyne_my_adb/internal/settings"
	"github.com/serhiq/go_fyne_my_adb/internal/term"
	"log"
)

const APP_UUID = "cd2f6c4e-6094-427a-b708-a4498eaa04d0"

type App struct {
	app             fyne.App
	win             fyne.Window
	MyTerm          *term.MyTerm
	cfg             config.Config
	runtimeSettings settings.RuntimeSettings
	cmds            []*commands.Cmd
	packageNameList map[string]struct{}
	cmdbar          *fyne.Container
	sidebar         *fyne.Container
}

func New(cfg config.Config) (*App, error) {

	packageNameList := map[string]struct{}{}

	for _, application := range cfg.Apps {
		packageNameList[application.PackageName] = struct{}{}
	}

	s, err := settings.Read(cfg.Main.Path)

	if err != nil {
		return nil, err
	}

	var app = &App{
		cfg:             cfg,
		packageNameList: packageNameList,
		cmds:            []*commands.Cmd{},
		runtimeSettings: *s,
	}

	return app, nil
}

func (s *App) Start() error {
	log.Println("app is starting...")

	s.app = app.NewWithID(APP_UUID)
	s.app.Settings().SetTheme(theme.DarkTheme())

	s.win = s.app.NewWindow(s.cfg.Main.Name)
	s.win.Resize(fyne.NewSize(800, 600))

	s.win.SetContent(s.MainScreen())

	s.win.SetCloseIntercept(func() {
		err := s.saveRuntimeSettings()
		if err != nil {
			log.Println(err)
		}
		s.app.Quit()

	})

	for _, cmd := range s.cmds {
		s.AddCmd(cmd)
	}

	s.win.ShowAndRun()

	return nil
}

func (s *App) Stop() {
	err := s.saveRuntimeSettings()
	if err != nil {
		println(err)
	}

	s.MyTerm.Exit()
	s.app.Quit()
}

func (s *App) saveRuntimeSettings() error {
	return s.runtimeSettings.SaveToFile(s.cfg.Main.Path)
}
