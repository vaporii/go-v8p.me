package domain

// struct representing a storage reference.
type StorageRef struct {
	// how the file is stored. currently only supports 'local'
	Provider string
	// provider-dependent string value. if local, stored as a path reference.
	Key string
}
