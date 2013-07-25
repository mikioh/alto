// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

import "encoding/json"

const (
	MediaTypeCostMap       = "application/alto-costmap+json"       // media type for ALTO map service
	MediaTypeCostMapFilter = "application/alto-costmapfilter+json" // media type for ALTO map filtering service
)

// A CostMap reprensents a list of path costs for each pair of
// source/destination provider-defined identifer (PID).
type CostMap struct {
	CostType   CostType            `json:"cost-type"`
	VersionTag string              `json:"map-vtag"`
	Map        map[string]DstCosts `json:"map"`
}

// MarshalJSON implements the MarshalJSON method of json.Marshaler
// interface.
func (cm *CostMap) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})
	raw["cost-type"] = cm.CostType
	raw["map-vtag"] = cm.VersionTag
	cmd := make(map[string]DstCosts)
	for pid, v := range cm.Map {
		cmd[pid] = v
	}
	raw["map"] = cmd
	return json.Marshal(raw)
}

// UnmarshalJSON implements the UnmarshalJSON method of
// json.Unmarshaler interface.
func (cm *CostMap) UnmarshalJSON(b []byte) error {
	var raw interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	for key, v := range raw.(map[string]interface{}) {
		switch key {
		case "cost-type":
			for key, vv := range v.(map[string]interface{}) {
				switch key {
				case "cost-metric":
					if v, ok := vv.(string); ok {
						cm.CostType.CostMetric = v
					}
				case "cost-mode":
					if v, ok := vv.(string); ok {
						cm.CostType.CostMode = v
					}
				case "description":
					if v, ok := vv.(string); ok {
						cm.CostType.Description = v
					}
				}
			}
		case "map-vtag":
			if v, ok := v.(string); ok {
				cm.VersionTag = v
			}
		case "map":
			cm.Map = make(map[string]DstCosts)
			for pid, vv := range v.(map[string]interface{}) {
				switch vv := vv.(type) {
				case map[string]interface{}:
					dcs := make(DstCosts)
					for pid, v := range vv {
						if v, ok := v.(float64); ok {
							dcs[pid] = v
						}
					}
					if len(dcs) > 0 {
						cm.Map[pid] = dcs
					}
				}
			}
		}
	}
	return nil
}

func (cm *CostMap) resourceType() string {
	return "costmap"
}

// A DstCosts represents a set of costs for the destination
// provider-defined identifier (PID).
type DstCosts map[string]float64

// A CostType represents a combination of cost type and cost mode.
type CostType struct {
	CostMetric  string `json:"cost-metric"`
	CostMode    string `json:"cost-mode"`
	Description string `json:"description,omitempty"`
}

// A ReqFilteredCostMap represents input parameters for the filtered
// cost map.
type ReqFilteredCostMap struct {
	CostType    CostType `json:"cost-type"`
	Constraints []string `json:"constraints,omitempty"`
	PIDs        struct {
		Srcs []string `json:"srcs,omitempty"`
		Dsts []string `json:"dsts,omitempty"`
	} `json:"pids,omitempty"`
}

// A FilteredCostMapCapabilities represents a capabilities for the
// filtered cost map.
type FilteredCostMapCapabilities struct {
	CostTypeNames   []string `json:"cost-type-names"`
	CostConstraints bool     `json:"cost-constraints"`
}
