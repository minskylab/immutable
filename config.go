package immutable

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type GotenbergConfig struct {
	Hostname string `yaml:"hostname"`
}

type PinataConfig struct {
	APIKey    string `yaml:"APIKey"`
	APISecret string `yaml:"APISecret"`
	// JWT       string `yaml:"JWT"`
}

type DocumentConfig struct {
	Title          string `yaml:"title"`
	OutputFilename string `yaml:"outputFilename"`
}

type Config struct {
	Document DocumentConfig `yaml:"document"`

	ImmutableDir string `yaml:"immutableDir"`
	TemplatesDir string `yaml:"templatesDir"`

	Gotenberg GotenbergConfig `yaml:"gotenberg"`
	Pinata    PinataConfig    `yaml:"pinata"`

	Production bool `yaml:"production"`
}

func LoadConfigFromFile(configPath string) (*Config, error) {
	config := Config{
		Document: DocumentConfig{
			Title:          "Immutable Document",
			OutputFilename: "document.pdf",
		},
		ImmutableDir: ".immutable",
		TemplatesDir: "document",
		Production:   false,
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	return &config, err
}
