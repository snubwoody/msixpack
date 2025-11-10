package bundle

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	Package     ConfigPackage
	Application ConfigApplication
}

type ConfigPackage struct {
	Version       string
	Name          string
	DisplayName   string `toml:"display-name"`
	Publisher     string
	PublisherName string `toml:"publisher-name"`
	Logo          string
	// Resources to include in the msix bundle.
	Resources string
}

type ConfigApplication struct {
	Id          string
	Executable  string
	Name        string
	Description string
}

func ReadConfig(path string) (Config, error) {
	var cfg Config
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = toml.Unmarshal(bytes, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
