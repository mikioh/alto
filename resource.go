// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

// A Resource represents an information resource.
type Resource struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

// A Meta represents a set of definitions related with the information
// resources.
type Meta map[string]interface{}

// A Data represents an information resource data.
type Data interface {
	resourceType() string
}

// NewResource returns an information resource. Known information
// resource types are "networkmap" and "costmap".
func NewResource(typ string) *Resource {
	switch typ {
	case "networkmap":
		return &Resource{Data: &NetworkMap{Map: make(map[string]EndpointAddrGroup)}}
	case "costmap":
		return &Resource{Data: &CostMap{Map: make(map[string]DstCosts)}}
	default:
		return &Resource{}
	}
}
