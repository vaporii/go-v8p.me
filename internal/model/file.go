package model

import "time"

// enum for different display types.
type DisplayAs string

const (
	// file will be represented as simple file info, and information will be embedded.
	FileInfo DisplayAs = "fileinfo"
	// file will be represented as a video, and embedded.
	Video DisplayAs = "video"
	// file will be represented as an image, and embedded.
	Image DisplayAs = "image"
	// file will be represented as text or code, and a preview will be embedded.
	Text DisplayAs = "text"
	// file will be represented as audio, and embedded where applicable.
	Audio DisplayAs = "audio"
	// file will be represented as markdown with an option to switch to raw view.
	// a preview will be embedded.
	Markdown DisplayAs = "markdown"
)

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
	// hidden from client for security reasons. stored relative to cwd.
	ServerPath string `json:"-"`
	// alias of the file.
	Alias string `json:"alias"`
	// domain name the file is hosted on.
	Domain string `json:"domain"`
	// date that the file will be deleted from the server and DB.
	// if the file does not expire, nil.
	Expires *time.Time `json:"expires,omitempty"`
	// how the file will be shown to the client. video, image, text, audio, or markdown.
	DisplayAs DisplayAs `json:"display_as"`
	// if a text file is recognized as a programming language, this is the highlighting.js
	// language to be used for syntax highlighting. nil if DisplayAs is not Text.
	// if the text is not recognized as a programming language, "text".
	HighlightingLanguage *string `json:"highlighting_language,omitempty"`
	// datetime the file was uploaded to the DB.
	CreatedAt time.Time `json:"created_at"`
}
