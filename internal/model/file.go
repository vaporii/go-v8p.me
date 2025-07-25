package model

import "time"

// struct representing metadata on a file.
type File struct {
	// ID representing the file in DB.
	ID string `json:"id"`
	// original name of the file including the file extension (if exists).
	Name string `json:"name"`
	// MIME type of the file. determined by client.
	Type string `json:"type"`
	// whether the file was encrypted.
	Encrypted bool `json:"encrypted"`
	// size of the file, in bytes.
	Size int64 `json:"size"`
	// randomly generated path of the file data on the server.
	// hidden from client for security reasons.
	ServerPath string `json:"-"`
	// alias of the file.
	Alias string `json:"alias"`
	// date that the file will be deleted from the server and DB.
	Expires time.Time `json:"expires"`
	// datetime the file was uploaded to the DB.
	CreatedAt time.Time `json:"created_at"`
}
