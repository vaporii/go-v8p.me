package config

import "os"

// global configuration. contains values used throughout the program.
type Config struct {
	// path to the SQLite database file. defaults to ./data/dev.db.
	SQLitePath string
	// path to a directory in which files will be stored. defaults to
	// ./data/files.
	FilesPath string
	// current environment. 'dev' or 'prod'. defaults to dev.
	Environment string
	// how detailed the logs will be. 'DEBUG', 'INFO', 'WARN', or 'ERROR'.
	// defaults to DEBUG.
	LogLevel string
	// the address the server will listen on. to listen on 0.0.0.0:3000,
	// you can use the shorthand ':3000'. defaults to ':3000'.
	ServerAddress string
}

var configuration *Config

// load the configuration. on first load, populates a local variable with the configuration
// and loads that on further calls.
func LoadConfig() *Config {
	if configuration == nil {
		configuration = &Config{
			SQLitePath:    loadEnv("V8P_SQLITE_PATH", "./data/dev.db"),
			FilesPath:     loadEnv("V8P_FILES_PATH", "./data/files"),
			Environment:   loadEnv("V8P_ENVIRONMENT", "dev"),
			LogLevel:      loadEnv("V8P_LOG_LEVEL", "DEBUG"),
			ServerAddress: loadEnv("V8P_SERVER_ADDRESS", ":3000"),
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
