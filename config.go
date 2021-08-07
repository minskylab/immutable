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
	JWT       string `yaml:"JWT"`
}

type Config struct {
	ImmutableDir string          `yaml:"immutableDir"`
	TemplatesDir string          `yaml:"templatesDir"`
	Gotenberg    GotenbergConfig `yaml:"gotenberg"`
	Pinata       PinataConfig    `yaml:"pinata"`
}

func LoadConfigFromFile(configPath string) (*Config, error) {
	config := Config{
		ImmutableDir: ".immutable",
		TemplatesDir: "document",
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	return &config, err
}
