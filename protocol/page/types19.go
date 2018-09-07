// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/protocol/network"
)

// FrameID Unique frame identifier.
//
// Provided as an alias to prevent circular dependencies.
type FrameID = internal.PageFrameID

// FrameID Unique frame identifier.
//type FrameID string

// Frame Information about the Frame on the page.
type Frame struct {
	ID             FrameID          `json:"id"`                 // Frame unique identifier.
	ParentID       *FrameID         `json:"parentId,omitempty"` // Parent frame identifier.
	LoaderID       network.LoaderID `json:"loaderId"`           // Identifier of the loader associated with this frame.
	Name           *string          `json:"name,omitempty"`     // Frame's name as specified in the tag.
	URL            string           `json:"url"`                // Frame document's URL.
	SecurityOrigin string           `json:"securityOrigin"`     // Frame document's security origin.
	MimeType       string           `json:"mimeType"`           // Frame document's mimeType as determined by the browser.
	// UnreachableURL If the frame failed to load, this contains the URL
	// that could not be loaded.
	//
	// Note: This property is experimental.
	UnreachableURL *string `json:"unreachableUrl,omitempty"`
}
