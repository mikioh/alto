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
	CostType    CostType `json:"cost-type"`
	Constraints []string `json:"constraints,omitempty"`
	Endpoints   struct {
		Srcs []Endpoint `json:"srcs,omitempty"`
		Dsts []Endpoint `json:"dsts,omitempty"`
	} `json:"endpoints"`
}

// An EndpointCostMap reprensents a list of endpoint cost maps.
type EndpointCostMap struct {
	CostType CostType                    `json:"cost-type"`
	Map      map[string]EndpointDstCosts `json:"map"`
}

// An EndpointDstCosts represents a set of endpoint cost maps.
type EndpointDstCosts map[string]interface{}
