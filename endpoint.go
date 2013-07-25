// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

import (
	"encoding/json"
	"net"
	"strings"

	"github.com/mikioh/ipaddr"
)

// An Endpoint represents an endpoint address or address prefix.
type Endpoint interface {
	Network() string
	String() string
	TypedString() string
}

// ParseEndpoint parses addr as a network endpoint identifier with
// address type typ. Known types are "ipv4" and "ipv6".
func ParseEndpoint(typ, addr string) (Endpoint, error) {
	switch typ {
	case "ipv4", "ipv6":
		return parseIPEndpoint(addr)
	case "mac-48", "mac-64":
		return parseMACEndpoint(addr)
	default:
		return nil, errUnknownAddress
	}
}

func splitTypedAddr(s string) (string, string) {
	i := strings.Index(s, ":")
	if i < 0 {
		return "", s
	}
	switch p := s[:i]; p {
	case "ipv4", "ipv6", "mac-48", "mac-64":
		return p, s[i+1:]
	}
	return "", s
}

// An IPEndpoint represents an IP address or address prefix.
type IPEndpoint struct {
	IP ipaddr.Prefix
}

// Network returns the endpoint's network; "ipv4" or "ipv6".
func (ep *IPEndpoint) Network() string {
	if ep == nil {
		return "<nil>"
	}
	if _, ok := ep.IP.(*ipaddr.IPv4); ok {
		return "ipv4"
	} else if _, ok := ep.IP.(*ipaddr.IPv6); ok {
		return "ipv6"
	}
	return "<nil>"
}

func (ep *IPEndpoint) String() string {
	if ep == nil {
		return "<nil>"
	}
	if _, ok := ep.IP.(*ipaddr.IPv4); ok && ep.IP.Len() != ipaddr.IPv4PrefixLen {
		return ep.IP.String()
	} else if _, ok := ep.IP.(*ipaddr.IPv6); ok && ep.IP.Len() != ipaddr.IPv6PrefixLen {
		return ep.IP.String()
	}
	return ep.IP.Addr().String()
}

// TypedString returns the literal endpoint address with network
// prefix followed by a colon.
func (ep *IPEndpoint) TypedString() string {
	if ep == nil {
		return "<nil>"
	}
	return ep.Network() + ":" + ep.String()
}

func parseIPEndpoint(s string) (ep *IPEndpoint, err error) {
	_, addr := splitTypedAddr(s)
	if ip := net.ParseIP(addr); ip != nil {
		if ipv4 := ip.To4(); ipv4 != nil {
			p, err := ipaddr.NewPrefix(ipv4, ipaddr.IPv4PrefixLen)
			if err != nil {
				return nil, err
			}
			ep = &IPEndpoint{IP: p}
		} else if ipv6 := ip.To16(); ipv6 != nil {
			p, err := ipaddr.NewPrefix(ipv6, ipaddr.IPv6PrefixLen)
			if err != nil {
				return nil, err
			}
			ep = &IPEndpoint{IP: p}
		} else {
			return nil, errUnknownAddress
		}
		return ep, nil
	}
	_, ipn, err := net.ParseCIDR(addr)
	if err != nil {
		return nil, err
	}
	l, _ := ipn.Mask.Size()
	p, err := ipaddr.NewPrefix(ipn.IP, l)
	if err != nil {
		return nil, err
	}
	return &IPEndpoint{IP: p}, nil
}

// A MACEndpoint represents a MAC address. Note that this address type
// is not defined in the ALTO protocol.
type MACEndpoint net.HardwareAddr

// Network returns the endpoint's network; "mac-48" or "mac-64".
func (ep MACEndpoint) Network() string {
	if len(ep) == 6 {
		return "mac-48"
	} else if len(ep) == 8 {
		return "mac-64"
	}
	return "<nil>"
}

func (ep MACEndpoint) String() string {
	return net.HardwareAddr(ep).String()
}

// TypedString returns the literal endpoint address with network
// prefix followed by a colon.
func (ep MACEndpoint) TypedString() string {
	return ep.Network() + ":" + ep.String()
}

func parseMACEndpoint(s string) (MACEndpoint, error) {
	_, addr := splitTypedAddr(s)
	hwa, err := net.ParseMAC(addr)
	if err != nil {
		return nil, err
	}
	return MACEndpoint(hwa), nil
}

// An EndpointAddrGroup represents a set of endpoints.
type EndpointAddrGroup map[string][]Endpoint

// MarshalJSON implements the MarshalJSON method of json.Marshaler
// interface.
func (eag EndpointAddrGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(eag.encode())
}

func (eag EndpointAddrGroup) encode() interface{} {
	raw := make(map[string][]string)
	for typ, v := range eag {
		ss := make([]string, len(v))
		for i := range v {
			ss[i] = v[i].String()
		}
		raw[typ] = ss
	}
	return raw
}

// UnmarshalJSON implements the UnmarshalJSON method of
// json.Unmarshaler interface.
func (eag EndpointAddrGroup) UnmarshalJSON(b []byte) error {
	var raw interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	return eag.decode(raw)
}

func (eag EndpointAddrGroup) decode(raw interface{}) error {
	for typ, v := range raw.(map[string]interface{}) {
		switch v := v.(type) {
		case []interface{}:
			var eps []Endpoint
			for _, e := range v {
				s, ok := e.(string)
				if !ok {
					continue
				}
				if ep, err := ParseEndpoint(typ, s); err != nil {
					return err
				} else {
					eps = append(eps, ep)
				}
			}
			if len(eps) > 0 {
				eag[typ] = eps
			}
		}
	}
	return nil
}

func (eag EndpointAddrGroup) toSlice(typ string) []Endpoint {
	if typ != "" {
		if eps, ok := eag[typ]; !ok {
			return nil
		} else {
			return eps
		}
	}
	var rs []Endpoint
	for _, eps := range eag {
		rs = append(rs, eps...)
	}
	return rs

}
