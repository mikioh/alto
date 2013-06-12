// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestDecodeEncodeNetworkMap(t *testing.T) {
	f, err := os.Open("testdata/networkmap.js")
	if err != nil {
		t.Fatalf("os.Open failed: %v", err)
	}
	nm := NetworkMap{}
	if err := json.NewDecoder(f).Decode(&nm); err != nil {
		t.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	if err := json.NewEncoder(&dst).Encode(&nm); err != nil {
		t.Fatalf("json.Encoder.Encode failed: %v", err)
	}
	var out bytes.Buffer
	if err := json.Indent(&out, dst.Bytes(), "", jsonIndent); err != nil {
		t.Fatalf("json.Indent failed: %v", err)
	} else {
		//t.Logf("%v", string(out.Bytes()))
	}
}

func TestDecodeEncodeNetworkMapData(t *testing.T) {
	f, err := os.Open("testdata/networkmapdata.js")
	if err != nil {
		t.Fatalf("os.Open failed: %v", err)
	}
	nmd := make(NetworkMapData)
	if err := json.NewDecoder(f).Decode(&nmd); err != nil {
		t.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	if err := json.NewEncoder(&dst).Encode(&nmd); err != nil {
		t.Fatalf("json.Encoder.Encode failed: %v", err)
	}
}

var networkMapTest = []struct {
	pid  string
	net  string
	addr string
}{
	{"pid1", "ipv4", "192.168.0.1"},
	{"pid1", "ipv4", "192.168.0.0/24"},
	{"pid1", "ipv6", "2001::1"},
	{"pid2", "ipv6", "2001::/32"},
}

func TestNetworkMap(t *testing.T) {
	nm := NewResource("networkmap").Data.(*NetworkMap)
	for _, tt := range networkMapTest {
		ep, err := ParseEndpoint(tt.net, tt.addr)
		if err != nil {
			t.Fatalf("ParseEndpoint failed: %v", err)
		}
		nm.Set(tt.pid, ep)
	}
	if eps := nm.Endpoints("", ""); len(eps) != 4 {
		t.Fatalf("got %v; expected %v", len(eps), 4)
	}
	if eps := nm.Endpoints("pid1", ""); len(eps) != 3 {
		t.Fatalf("got %v; expected %v", len(eps), 3)
	}
	if eps := nm.Endpoints("pid1", "ipv4"); len(eps) != 2 {
		t.Fatalf("got %v; expected %v", len(eps), 2)
	}
	if eps := nm.Endpoints("pid1", "ipv6"); len(eps) != 1 {
		t.Fatalf("got %v; expected %v", len(eps), 1)
	}
	if eps := nm.Endpoints("pid2", ""); len(eps) != 1 {
		t.Fatalf("got %v; expected %v", len(eps), 1)
	}
	if eps := nm.Endpoints("pid2", "ipv4"); len(eps) != 0 {
		t.Fatalf("got %v; expected %v", len(eps), 0)
	}
	if eps := nm.Endpoints("pid2", "ipv6"); len(eps) != 1 {
		t.Fatalf("got %v; expected %v", len(eps), 1)
	}
}
