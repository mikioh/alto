// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

import "encoding/json"

const (
	MediaTypeNetworkMap       = "application/alto-networkmap+json"       // media type for ALTO map service
	MediaTypeNetworkMapFilter = "application/alto-networkmapfilter+json" // media type for ALTO map filtering service
)

// A NetworkMap represents a list of network locations within the
// provider-defined identifier (PID).
type NetworkMap struct {
	VersionTag string                       `json:"map-vtag"`
	Map        map[string]EndpointAddrGroup `json:"map"`
}

// MarshalJSON implements the MarshalJSON method of json.Marshaler
// interface.
func (nm *NetworkMap) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})
	raw["map-vtag"] = nm.VersionTag
	nmd := make(map[string]interface{})
	for pid, v := range nm.Map {
		nmd[pid] = v.encode()
	}
	raw["map"] = nmd
	return json.Marshal(raw)
}

// UnmarshalJSON implements the UnmarshalJSON method of
// json.Unmarshaler interface.
func (nm *NetworkMap) UnmarshalJSON(b []byte) error {
	var raw interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	for key, v := range raw.(map[string]interface{}) {
		switch key {
		case "map-vtag":
			if v, ok := v.(string); ok {
				nm.VersionTag = v
			}
		case "map":
			if len(nm.Map) == 0 {
				nm.Map = make(map[string]EndpointAddrGroup)
			}
			for pid, vv := range v.(map[string]interface{}) {
				switch vv := vv.(type) {
				case map[string]interface{}:
					eag := make(EndpointAddrGroup)
					if err := eag.decode(vv); err != nil {
						return err
					}
					nm.Map[pid] = eag
				}
			}
		}
	}
	return nil
}

func (nm *NetworkMap) resourceType() string {
	return "networkmap"
}

// Endpoints returns a list of endpoints which selected with
// provider-defined identifier (PID) name pid and address type
// typ. The zero value for string is treated as wildcard.
func (nm *NetworkMap) Endpoints(pid, typ string) []Endpoint {
	eags := nm.toSlice(pid)
	if eags == nil {
		return nil
	}
	if len(eags) == 1 {
		return eags[0].toSlice(typ)
	}
	var eps []Endpoint
	for _, eag := range eags {
		eps = append(eps, eag.toSlice(typ)...)
	}
	return eps
}

func (nm *NetworkMap) toSlice(pid string) []EndpointAddrGroup {
	if pid != "" {
		if eag, ok := nm.Map[pid]; !ok {
			return nil
		} else {
			return []EndpointAddrGroup{eag}
		}
	}
	var eags []EndpointAddrGroup
	for _, eag := range nm.Map {
		eags = append(eags, eag)
	}
	return eags
}

// Set sets the provider-defined identifier (PID) pid to endpoiint
// ep. It replaces any existing endpoints.
func (nm *NetworkMap) Set(pid string, ep Endpoint) {
	if eag, ok := nm.Map[pid]; !ok {
		eag := make(EndpointAddrGroup)
		eag[ep.Network()] = []Endpoint{ep}
		nm.Map[pid] = eag
	} else {
		if eps, ok := eag[ep.Network()]; !ok {
			eag[ep.Network()] = []Endpoint{ep}
		} else {
			eps = append(eps, ep)
			eag[ep.Network()] = eps
		}
		nm.Map[pid] = eag
	}
}

// A ReqFilteredNetworkMap represents input parameters for the
// filtered network map.
type ReqFilteredNetworkMap struct {
	PIDs      []string `json:"pids,omitempty"`
	AddrTypes []string `json:"address-types,omitempty"`
}
