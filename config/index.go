package config

import (
	"os"

	"github.com/rs/zerolog"
)

func GetKey(key string) string {
	return os.Getenv(key)
}

type GlobalConfig struct {
	Dotfile string
}

var cfg = &GlobalConfig{}

func GetLogger() zerolog.Logger {
	logger := zerolog.New(os.Stdout).
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	return logger
}

func GetGlobalConfig() *GlobalConfig {
	return cfg
}
