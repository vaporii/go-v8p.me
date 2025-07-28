package dto

import "time"

// struct representing the data that will be sent from the client.
type CreateFileRequest struct {
	// filename of the file
	Name string `json:"name"`
	// MIME type of the file. e.g. 'application/json; charset=utf-8'
	Type string `json:"type"`
	// whether the file data is encrypted.
	Encrypted bool `json:"encrypted"`
	// size of the file, in bytes.
	Size int64 `json:"size"`
	// file expiration date. nil if none.
	Expires *time.Time `json:"expires,omitempty"`
	// what the file will display as.
	// "fileinfo", "video", "image", "text", "audio", or "markdown".
	DisplayAs string `json:"display_as"`
	// if DisplayAs=="text", will be "text" or one of many different
	// syntax highlighting languages.
	HighlightingLanguage *string `json:"highlighting_language,omitempty"`
}
