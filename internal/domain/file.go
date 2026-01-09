package domain

import "time"

// struct representing metadata on a file.
type File struct {
	// ID representing the file in DB.
	ID int64 `json:"id"`
	// original name of the file including the file extension (if exists).
	Name string `json:"name"`
	// MIME type of the file. determined by client.
	Type string `json:"type"`
	// whether the file was encrypted.
	Encrypted bool `json:"encrypted"`
	// size of the file, in bytes.
	Size int64 `json:"size"`
	// alias of the file.
	Alias string `json:"alias"`
	// date that the file will be deleted from the server and DB.
	// if the file does not expire, nil.
	Expires *time.Time `json:"expires,omitempty"`
	// if a text file is recognized as a programming language, this is the highlighting.js
	// language to be used for syntax highlighting. nil if not text.
	// if the text is not recognized as a programming language, "text".
	HighlightingLanguage *string `json:"highlighting_language,omitempty"`
	// how the file is stored.
	Storage StorageRef
	// datetime the file was uploaded to the DB.
	CreatedAt time.Time `json:"created_at"`
}
