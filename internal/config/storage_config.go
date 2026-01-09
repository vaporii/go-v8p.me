package config

type LocalStorageConfig struct {
	// path to a directory in which files will be stored. defaults to
	// ./data/files.
	BaseFilesPath string
	// character length of the randomly generated filenames
	FileNameLength int
}
