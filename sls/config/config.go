package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Listener []string       `json:"listener"`
	Sender   map[string]any `json:"sender"`
}

func LoadConfig(files []string) (*Config, error) {
	var (
		cfg *Config
		erf []string
	)
	for _, file := range files {
		jf, err := os.ReadFile(file)
		if err != nil {
			erf = append(erf, file)
			continue
		}

		var c Config
		if err := json.Unmarshal(jf, &c); err != nil {
			erf = append(erf, file)
			continue
		}
		cfg = &c
		break
	}
	if cfg == nil {
		return nil, fmt.Errorf("config file not found %s", strings.Join(erf, ","))
	}
	return cfg, nil
}
