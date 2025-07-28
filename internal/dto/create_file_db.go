package dto

import "time"

// struct representing the data that will be sent from the service to the repository.
type CreateFileDB struct {
	// original name of the file including the file extension (if exists).
	Name string `json:"name"`
	// MIME type of the file. determined by client.
	Type string `json:"type"`
	// whether the file was encrypted.
	Encrypted bool `json:"encrypted"`
	// size of the file, in bytes.
	Size int64 `json:"size"`
	// randomly generated path of the file data on the server.
	// hidden from client for security reasons. stored relative to cwd.
	ServerPath string `json:"-"`
	// alias of the file.
	Alias string `json:"alias"`
	// date that the file will be deleted from the server and DB.
	// if the file does not expire, nil.
	Expires *time.Time `json:"expires,omitempty"`
	// how the file will be shown to the client. video, image, text, audio, or markdown.
	DisplayAs string `json:"display_as"`
	// if a text file is recognized as a programming language, this is the highlighting.js
	// language to be used for syntax highlighting. nil if DisplayAs is not Text.
	// if the text is not recognized as a programming language, "text".
	HighlightingLanguage *string `json:"highlighting_language,omitempty"`
}
