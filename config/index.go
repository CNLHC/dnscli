package config

import "os"

func GetKey(key string) string {
	return os.Getenv(key)
}

type GlobalConfig struct {
	Dotfile string
}

var cfg = &GlobalConfig{}

func GetGlobalConfig() *GlobalConfig {
	return cfg
}
