package login

import (
	"flag"
	"pandora/internal/pkg"
)

type Config struct {
	Env          pkg.Environment
	LogLevel     string
	DebugMode    bool
	ServerPort   int
	DatabaseUri  string
	DatabaseName string
}

func NewConfig() (*Config, error) {
	config := &Config{
		Env:          pkg.Environment(pkg.GetEnv("ENV", "local")),
		LogLevel:     pkg.GetEnv("LOG_LEVEL", "info"),
		DebugMode:    pkg.GetEnvAsBool("DEBUG_MODE", false),
		ServerPort:   pkg.GetEnvAsInt("SERVER_PORT", 443),
		DatabaseUri:  pkg.GetEnv("DATABASE_URI", "mongodb://localhost:27017"),
		DatabaseName: pkg.GetEnv("DATABASE_NAME", "pandora"),
	}

	flag.IntVar(&config.ServerPort, "port", config.ServerPort, "port to listen on")
	flag.BoolVar(&config.DebugMode, "debug", config.DebugMode, "enable debug mode")
	flag.Parse()

	return config, nil
}
