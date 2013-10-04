// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

import (
	"bytes"
	"encoding/json"
	"net"
	"os"
	"reflect"
	"testing"
)

func TestDecodeEncodeEndpointAddrGroup(t *testing.T) {
	f, err := os.Open("testdata/endpointaddrgroup.js")
	if err != nil {
		t.Fatalf("os.Open failed: %v", err)
	}
	eag := make(EndpointAddrGroup)
	if err := json.NewDecoder(f).Decode(&eag); err != nil {
		t.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	if err := json.NewEncoder(&dst).Encode(&eag); err != nil {
		t.Fatalf("json.Encoder.Encode failed: %v", err)
	}
}

var parseIPEndpointTests = []struct {
	net       string
	in        string
	ip        net.IP
	prefixLen int
}{
	{"ipv4", "172.16.254.191", net.ParseIP("172.16.254.191"), 32},
	{"ipv4", "172.16.254.192/32", net.ParseIP("172.16.254.192"), 32},
	{"ipv4", "ipv4:172.16.254.193", net.ParseIP("172.16.254.193"), 32},
	{"ipv4", "172.16.254.194/24", net.ParseIP("172.16.254.0"), 24},
	{"ipv4", "ipv4:172.16.254.195/24", net.ParseIP("172.16.254.0"), 24},
	{"ipv4", "ipv4:172.16.254.0/24", net.ParseIP("172.16.254.0"), 24},

	{"ipv6", "2001:abcd::191", net.ParseIP("2001:abcd::191"), 128},
	{"ipv6", "2001:abcd::192/128", net.ParseIP("2001:abcd::192"), 128},
	{"ipv6", "ipv6:2001:abcd::193/128", net.ParseIP("2001:abcd::193"), 128},
	{"ipv6", "2001:abcd::194/29", net.ParseIP("2001:abc8::"), 29},
	{"ipv6", "ipv6:2001:abcd::195/29", net.ParseIP("2001:abc8::"), 29},
	{"ipv6", "ipv6:2001:abcd::/29", net.ParseIP("2001:abc8::"), 29},
}

func TestParseIPEndpoint(t *testing.T) {
	for _, tt := range parseIPEndpointTests {
		ep, err := ParseEndpoint(tt.net, tt.in)
		if err != nil {
			t.Fatalf("ParseEndpoint(%q, %q) failed: %v", tt.net, tt.in, err)
		}
		switch ep := ep.(type) {
		case *IPEndpoint:
			if !ep.IP.Addr().Equal(tt.ip) || ep.IP.Len() != tt.prefixLen {
				t.Fatalf("got %v; expected %v/%v", ep, tt.ip, tt.prefixLen)
			}
		default:
			t.Fatalf("got unknown endpoint %v", ep)
		}
	}
}

var parseMACEndpointTests = []struct {
	net string
	in  string
	out Endpoint
}{
	{"mac-48", "01:23:45:67:89:ab", MACEndpoint{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}},
	{"mac-48", "mac-48:01:23:45:67:89:ab", MACEndpoint{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}},

	{"mac-64", "01:23:45:67:89:ab:cd:ef", MACEndpoint{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}},
	{"mac-64", "mac-64:01:23:45:67:89:ab:cd:ef", MACEndpoint{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}},
}

func TestParseMACEndpoint(t *testing.T) {
	for _, tt := range parseMACEndpointTests {
		ep, err := ParseEndpoint(tt.net, tt.in)
		if err != nil {
			t.Fatalf("ParseEndpoint(%q, %q) failed: %v", tt.net, tt.in, err)
		}
		switch ep := ep.(type) {
		case MACEndpoint:
			if !reflect.DeepEqual(ep, tt.out) {
				t.Fatalf("got %v; expected %v", ep, tt.out)
			}
		default:
			t.Fatalf("got unknown endpoint %v", ep)
		}
	}
}
