// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

const (
	MediaTypeDirectory = "application/alto-directory+json" // media type for ALTO directory service
)

// A Directory represents an information resource directory.
type Directory struct {
	Meta      Meta                `json:"meta"`
	Resources []DirectoryResource `json:"resources"`
}

// A DirResource represents a list of information resources.
type DirectoryResource struct {
	URI          string                 `json:"uri"`
	MediaType    string                 `json:"media-type"`
	Accepts      string                 `json:"accepts,omitempty"`
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
}
