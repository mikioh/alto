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

func TestDecodeEncodeCostMap(t *testing.T) {
	f, err := os.Open("testdata/costmap.js")
	if err != nil {
		t.Fatalf("os.Open failed: %v", err)
	}
	cm := CostMap{}
	if err := json.NewDecoder(f).Decode(&cm); err != nil {
		t.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	if err := json.NewEncoder(&dst).Encode(&cm); err != nil {
		t.Fatalf("json.Encoder.Encode failed: %v", err)
	}
	var out bytes.Buffer
	if err := json.Indent(&out, dst.Bytes(), "", jsonIndent); err != nil {
		t.Fatalf("json.Indent failed: %v", err)
	} else {
		//t.Logf("%v", string(out.Bytes()))
	}
}

func TestDecodeEncodeCostMapData(t *testing.T) {
	f, err := os.Open("testdata/costmapdata.js")
	if err != nil {
		t.Fatalf("os.Open failed: %v", err)
	}
	cmd := make(CostMapData)
	if err := json.NewDecoder(f).Decode(&cmd); err != nil {
		t.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	if err := json.NewEncoder(&dst).Encode(&cmd); err != nil {
		t.Fatalf("json.Encoder.Encode failed: %v", err)
	}
}

func TestDecodeEncodeDstCosts(t *testing.T) {
	f, err := os.Open("testdata/dstcosts.js")
	if err != nil {
		t.Fatalf("os.Open failed: %v", err)
	}
	dcs := make(DstCosts)
	if err := json.NewDecoder(f).Decode(&dcs); err != nil {
		t.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	if err := json.NewEncoder(&dst).Encode(&dcs); err != nil {
		t.Fatalf("json.Encoder.Encode failed: %v", err)
	}
}
