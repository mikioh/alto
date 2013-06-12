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

func TestDecodeEncodeDirectory(t *testing.T) {
	f, err := os.Open("testdata/directory.js")
	if err != nil {
		t.Fatalf("os.Open failed: %v", err)
	}
	d := Directory{}
	if err := json.NewDecoder(f).Decode(&d); err != nil {
		t.Fatalf("json.Decoder.Decode failed: %v", err)
	}
	var dst bytes.Buffer
	if err := json.NewEncoder(&dst).Encode(&d); err != nil {
		t.Fatalf("json.Encoder.Encode failed: %v", err)
	}
	var out bytes.Buffer
	if err := json.Indent(&out, dst.Bytes(), "", jsonIndent); err != nil {
		t.Fatalf("json.Indent failed: %v", err)
	} else {
		t.Logf("%v", string(out.Bytes()))
	}
}
