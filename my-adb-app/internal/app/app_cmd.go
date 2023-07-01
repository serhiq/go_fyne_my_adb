package app

import (
	"fmt"
	"github.com/serhiq/go_fyne_my_adb/internal/commands"
	"github.com/serhiq/go_fyne_my_adb/internal/screenshot"
	"strings"
)

func (s *App) runScreencap() {
	filename, err := s.getScreenShotPathWithName()
	if err != nil {
		println(err)
	}

	var c = commands.Screencap(filename)
	s.sendCmd(c)
}

func (s *App) getScreenShotPathWithName() (string, error) {
	sb := strings.Builder{}

	if s.cfg.Screenshot.Path != "" {
		sb.WriteString(s.cfg.Screenshot.Path)

		//if s.runtimeSettings.ScreenshotFolder != "" {
		//	sb.WriteString("/")
		//	sb.WriteString(s.runtimeSettings.ScreenshotFolder)
		//}
		sb.WriteString("/")

		sb.WriteString(screenshot.GenerateFilename("screen", s.runtimeSettings.GetScreenshotCounter()))
		//sb.WriteString(GenerateFilename(s.runtimeSettings.ScreenshotName(), s.runtimeSettings.screenshotCounter()))

	} else {
		return "", fmt.Errorf("screenshot path not set")
	}

	return sb.String(), nil
}
