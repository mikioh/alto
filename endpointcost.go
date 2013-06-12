// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

const (
	MediaTypeEndpointCost       = "application/alto-endpointcost+json"       // media type for ALTO endpoint cost service
	MediaTypeEndpointCostParams = "application/alto-endpointcostparams+json" // media type for ALTO endpoint cost service
)

// A ReqEndpointCostMap represents input parameters for the filtered
// cost map.
type ReqEndpointCostMap struct {
	CostType    CostType       `json:"cost-type"`
	Constraints []string       `json:"constraints,omitempty"`
	Endpoints   EndpointFilter `json:"endpoints"`
}

// An EndpointFilter represents a list of source and destionation
// endpoints for which path costs are to be returned.
type EndpointFilter struct {
	Srcs []Endpoint `json:"srcs,omitempty"`
	Dsts []Endpoint `json:"dsts,omitempty"`
}

// An EndpointCostMap reprensents a list of endpoint cost maps.
type EndpointCostMap struct {
	CostType CostType            `json:"cost-type"`
	Map      EndpointCostMapData `json:"map"`
}

// An EndpointCostMapData represents a set of endpoint cost maps for
// the given destination endpoints.
type EndpointCostMapData map[string]EndpointDstCosts

// An EndpointDstCosts represents a set of endpoint cost maps.
type EndpointDstCosts map[string]interface{}
