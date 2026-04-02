package env

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultConfigPath = ".env"

// Parse tries to read envs from .env file firstly, then from application envs
// env is a pointer to the Config structure
func Parse(config any) error {
	configPath := os.Getenv("CONFIG")
	if configPath == "" {
		configPath = defaultConfigPath
	}

	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		if err := cleanenv.ReadEnv(config); err != nil {
			return fmt.Errorf("failed to read env vars: %w", err)
		}
	}

	return nil
}
