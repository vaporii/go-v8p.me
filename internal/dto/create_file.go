package dto

import "time"

type CreateFile struct {
	Name                 string     `json:"id"`
	Type                 string     `json:"type"`
	Encrypted            bool       `json:"encrypted"`
	Size                 int64      `json:"size"`
	Domain               string     `json:"domain"`
	Expires              *time.Time `json:"expires,omitempty"`
	DisplayAs            string     `json:"display_as"`
	HighlightingLanguage *string    `json:"highlighting_language,omitempty"`
}
