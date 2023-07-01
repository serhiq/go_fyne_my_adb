package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
)

type Config struct {
	Main       Main       `yaml:"app"`
	Commands   []Command  `yaml:"commands"`
	Apps       []Package  `yaml:"apps"`
	Screenshot Screenshot `yaml:"screenshot"`
}

type Main struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type Command struct {
	Cmd             string `yaml:"cmd"`
	Description     string `yaml:"description"`
	PassListPackage bool   `yaml:"passListPackage,omitempty"`
}

type Package struct {
	PackageName string `yaml:"packageName"`
}

type Screenshot struct {
	Path string `yaml:"path"`
}

func New() (*Config, error) {

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "./config/config.yaml"
	}
	config := &Config{}

	if err := fromYaml(path, config); err != nil {
		fmt.Printf("couldn'n load config from %s: %s\r\n", path, err.Error())
	}

	if config.Screenshot.Path != "" {
		if err := os.MkdirAll(filepath.Dir(config.Screenshot.Path), fs.ModeDir); err != nil {
			return nil, fmt.Errorf("config: failed creating  path %s (%s)", filepath.Dir(config.Screenshot.Path), err)
		}
	}

	return config, nil
}

func fromYaml(path string, config *Config) error {
	if path == "" {
		return nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, config)
}
