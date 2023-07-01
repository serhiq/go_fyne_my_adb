package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const APP_RUNTIMESETTINGSNAME = "settings.json"

type RuntimeSettings struct {
	ScreenshotFolder  string `json:"screenshotFolder"`
	ScreenshotCounter int    `json:"GetScreenshotCounter"`
}

func createPathSettings(path string) string {
	return path + "/" + APP_RUNTIMESETTINGSNAME
}
func (s *RuntimeSettings) GetScreenshotCounter() int {
	s.ScreenshotCounter = s.ScreenshotCounter + 1
	return s.ScreenshotCounter
}

func (s *RuntimeSettings) SaveToFile(path string) error {
	return saveSettingsToFile(createPathSettings(path), *s)
}

func saveSettingsToFile(filename string, settings RuntimeSettings) error {

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	settingsJSON, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %s", err.Error())
	}
	_, err = f.Write(settingsJSON)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %s", err.Error())
	}

	err = f.Sync() // Flush the changes to disk
	if err != nil {
		return fmt.Errorf("ошибка синхронизации файла: %s", err.Error())
	}

	f.Close()
	content, err := ioutil.ReadFile(filename)
	fmt.Printf("Содержимое файла: %s\n", content)
	println(content)

	wd, err := os.Getwd()
	println(wd)

	return nil
}

func Read(mainPath string) (*RuntimeSettings, error) {

	setting, err := readOrCreateJSONFile(mainPath, RuntimeSettings{
		ScreenshotFolder:  "",
		ScreenshotCounter: 1,
	})
	if err != nil {
		return nil, err
	}
	return setting, nil

}

func readOrCreateJSONFile(mainPath string, defaultSettings RuntimeSettings) (*RuntimeSettings, error) {
	var path = createPathSettings(mainPath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {

		err := defaultSettings.SaveToFile(path)
		if err != nil {
			return nil, err
		}

		return &defaultSettings, nil
	}

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %s", err.Error())
	}

	var s RuntimeSettings
	err = json.Unmarshal(fileContent, &s)
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования JSON: %s", err.Error())
	}

	return &s, nil
}
