package config

import "os"

// global configuration. contains values used throughout the program.
type Config struct {
	// path to the SQLite database file.
	SQLitePath string
	// current environment. 'dev' or 'prod'.
	Environment string
}

var configuration *Config

// load the configuration. on first load, populates a local variable with the configuration
// and loads that on further calls.
func LoadConfig() *Config {
	if configuration == nil {
		configuration = &Config{
			SQLitePath:  loadEnv("V8P_SQLITE_PATH", "./dev.db"),
			Environment: loadEnv("V8P_ENVIRONMENT", "dev"),
		}
	}

	return configuration
}

func loadEnv(envName string, defaultVal string) string {
	if value, exists := os.LookupEnv(envName); exists {
		return value
	}
	return defaultVal
}
