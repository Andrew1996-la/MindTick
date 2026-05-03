package logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func newConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("LOG", &config); err != nil {
		return Config{}, fmt.Errorf("env process error: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := newConfig()
	if err != nil {
		err = fmt.Errorf("get logger config: %w", err)
		panic(err)
	}

	return config
}
