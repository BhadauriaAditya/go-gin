package config

import "github.com/google/wire"

// ProvideConfig loads and returns the Config
func ProvideConfig() *Config {
	return LoadEnv()
}

var ConfigSet = wire.NewSet(ProvideConfig)