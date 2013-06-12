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

const jsonIndent = "    "

func BenchmarkDecodeNetworkMap(b *testing.B) {
	f, err := os.Open("testdata/resource-networkmap.js")
	if err != nil {
		b.Fatalf("os.Open failed: %v", err)
	}
	nm := NewResource("networkmap")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		if err := json.NewDecoder(f).Decode(nm); err != nil {
			b.Fatalf("json.Decoder.Decode failed: %v", err)
		}
	}
}

func BenchmarkEncodeNetworkMap(b *testing.B) {
	f, err := os.Open("testdata/resource-networkmap.js")
	if err != nil {
		b.Fatalf("os.Open failed: %v", err)
	}
	nm := NewResource("networkmap")
	if err := json.NewDecoder(f).Decode(nm); err != nil {
		b.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := json.NewEncoder(&dst).Encode(nm); err != nil {
			b.Fatalf("json.Encoder.Encode failed: %v", err)
		}
	}
}

var resourceTests = []struct {
	name string
	typ  string
}{
	{"testdata/resource-networkmap.js", "networkmap"},
	{"testdata/resource-costmap.js", "costmap"},
}

func TestDecodeEncodeResource(t *testing.T) {
	for _, tt := range resourceTests {
		f, err := os.Open(tt.name)
		if err != nil {
			t.Fatalf("os.Open failed: %v", err)
		}
		r := NewResource(tt.typ)
		if err := json.NewDecoder(f).Decode(r); err != nil {
			t.Fatalf("json.Decoder.Decode failed: %v", err)
		}
		var dst bytes.Buffer
		if err := json.NewEncoder(&dst).Encode(r); err != nil {
			t.Fatalf("json.Encoder.Encode failed: %v", err)
		}
		var out bytes.Buffer
		if err := json.Indent(&out, dst.Bytes(), "", jsonIndent); err != nil {
			t.Fatalf("json.Indent failed: %v", err)
		} else {
			t.Logf("%v", string(out.Bytes()))
		}
	}
}
