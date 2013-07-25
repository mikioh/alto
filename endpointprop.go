// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

const (
	MediaTypeEndpointProp       = "application/alto-endpointprop+json"       // media type for ALTO endpoint property service
	MediaTypeEndpointPropParams = "application/alto-endpointpropparams+json" // media type for ALTO endpoint property service
)

// A ReqEndpointProp represents input parameters for the filtered
// endpoint properties.
type ReqEndpointProp struct {
	Properties []string   `json:"properties"`
	Endpoints  []Endpoint `json:"endpoints"`
}

// An EndpointPropertyCapabilities reprensents a capabilities of
// endpoint property.
type EndpointPropertyCapabilities struct {
	PropTypes []string `json:"prop-types"`
}

// An EndpointProperty represents a list of endpoint properties.
type EndpointProperty struct {
	VersionTag string                   `json:"map-vtag"`
	Map        map[string]EndpointProps `json:"map"`
}

// An EndpointProps represents a set of endpoint properties.
type EndpointProps map[string]interface{}
