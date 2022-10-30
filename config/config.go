package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

func Config() *koanf.Koanf {
	osEnv := os.Getenv("INDEXER_ENV")
	configPath := os.Getenv("INDEXER_CONFIG_PATH")
	cfg := koanf.New(".")
	cfg.Load(file.Provider(fmt.Sprintf("%s/config.yaml", configPath)), yaml.Parser())
	cfg.Load(file.Provider(fmt.Sprintf("%s/config.%s.yaml", configPath, osEnv)), yaml.Parser())

	// Load environment variables and merge into the loaded config.
	// "INDEXER" is the prefix to filter the osEnv vars by.
	// "." is the delimiter used to represent the key hierarchy in osEnv vars.
	// The (optional, or can be nil) function can be used to transform
	// the osEnv var names, for instance, to lowercase them.
	//
	// For example, osEnv vars: INDEXER_TYPE and INDEXER_PARENT1_CHILD1_NAME
	// will be merged into the "type" and the nested "parent1.child1.name"
	// keys in the config file here as we lowercase the key,
	// replace `_` with `.` and strip the INDEXER_ prefix so that
	// only "parent1.child1.name" remains.
	cfg.Load(env.Provider("INDEXER_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "INDEXER_")), "_", ".", -1)
	}), nil)

	return cfg
}
