package config

import (
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func GetKey(key string) string {
	return viper.GetString(key)
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
func GetGoModRoot() string {
	cwd, e := os.Getwd()
	if e != nil {
		return "."
	}
	cur := cwd
	depth := 0
	for depth < 5 {
		stat, _ := os.Stat(path.Join(cur, "go.mod"))
		if stat != nil {
			return cur
		}
		cur = path.Join(cur, "..")
		depth += 1
	}
	return "."
}

func init() {
}
