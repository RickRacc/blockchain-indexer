package config

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"os"
)

func Config() *koanf.Koanf {
	env := os.Getenv("INDEXER_ENV")
	configPath := os.Getenv("CONFIG_PATH")
	cfg := koanf.New(".")
	cfg.Load(file.Provider(fmt.Sprintf("%s/config.yaml", configPath)), yaml.Parser())
	cfg.Load(file.Provider(fmt.Sprintf("%s/config.%s.yaml", configPath, env)), yaml.Parser())

	return cfg
}
