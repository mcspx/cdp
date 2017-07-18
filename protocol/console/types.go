// Code generated by cdpgen. DO NOT EDIT.

package console

// Message Console message.
type Message struct {
	Source string  `json:"source"`           // Message source.
	Level  string  `json:"level"`            // Message severity.
	Text   string  `json:"text"`             // Message text.
	URL    *string `json:"url,omitempty"`    // URL of the message origin.
	Line   *int    `json:"line,omitempty"`   // Line number in the resource that generated this message (1-based).
	Column *int    `json:"column,omitempty"` // Column number in the resource that generated this message (1-based).
}