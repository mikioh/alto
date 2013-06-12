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
	VersionTag string         `json:"map-vtag"`
	Map        NetworkMapData `json:"map"`
}

// MarshalJSON implements the MarshalJSON method of json.Marshaler
// interface.
func (nm *NetworkMap) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})
	raw["map-vtag"] = nm.VersionTag
	raw["map"] = nm.Map.encode()
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
			nm.Map = make(NetworkMapData)
			if err := nm.Map.decode(v); err != nil {
				return err
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
	eags := nm.Map.toSlice(pid)
	if eags == nil {
		return nil
	}
	if len(eags) == 1 {
		return eags[0].toSlice(typ)
	}
	var rs []Endpoint
	for _, eag := range eags {
		rs = append(rs, eag.toSlice(typ)...)
	}
	return rs
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

// A NetworkMapData represents a set of endpoint addresses within the
// provider-defined identifier (PID).
type NetworkMapData map[string]EndpointAddrGroup

// MarshalJSON implements the MarshalJSON method of json.Marshaler
// interface.
func (nmd NetworkMapData) MarshalJSON() ([]byte, error) {
	return json.Marshal(nmd.encode())
}

func (nmd NetworkMapData) encode() interface{} {
	raw := make(map[string]interface{})
	for pid, v := range nmd {
		raw[pid] = v.encode()
	}
	return raw
}

// UnmarshalJSON implements the UnmarshalJSON method of
// json.Unmarshaler interface.
func (nmd NetworkMapData) UnmarshalJSON(b []byte) error {
	var raw interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	return nmd.decode(raw)
}

func (nmd NetworkMapData) decode(raw interface{}) error {
	for pid, v := range raw.(map[string]interface{}) {
		switch v := v.(type) {
		case map[string]interface{}:
			eag := make(EndpointAddrGroup)
			if err := eag.decode(v); err != nil {
				return err
			}
			nmd[pid] = eag
		}
	}
	return nil
}

func (nmd NetworkMapData) toSlice(pid string) []EndpointAddrGroup {
	if pid != "" {
		if eag, ok := nmd[pid]; !ok {
			return nil
		} else {
			return []EndpointAddrGroup{eag}
		}
	}
	var rs []EndpointAddrGroup
	for _, eag := range nmd {
		rs = append(rs, eag)
	}
	return rs
}

// A ReqFilteredNetworkMap represents input parameters for the
// filtered network map.
type ReqFilteredNetworkMap struct {
	PIDs      []string `json:"pids,omitempty"`
	AddrTypes []string `json:"address-types,omitempty"`
}
