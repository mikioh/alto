// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

// +build ignore

package altoserver

import (
	"appengine"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mikioh/alto"
	"net"
	"net/http"
)

var (
	nmap *alto.Resource
)

func init() {
	http.HandleFunc("/", handler)
	nmap = alto.NewResource("networkmap")
	json.Unmarshal(data, nmap)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var bb bytes.Buffer
	if err := json.NewEncoder(&bb).Encode(nmap); err != nil {
		c.Errorf("json.Encoder.Encode failed: %v", err)
		return
	}
	bb.Reset()
	fmt.Fprintf(&bb, "%v, %v, %v, %v", r.Header.Get("X-AppEngine-Country"), r.Header.Get("X-AppEngine-Region"), r.Header.Get("X-AppEngine-City"), r.Header.Get("X-AppEngine-CityLatLong"))
	w.Header().Set("Content-Type", alto.MediaTypeNetworkMap)
	fmt.Fprintf(&bb, "Summary: %s", sum.String())
	if _, err := w.Write(bb.Bytes()); err != nil {
		c.Errorf("http.ResponseWriter.Write failed: %v", err)
		return
	}
}

var data = []byte(`{ "meta": { "redistribution": { "service-id": "12ab34cd", "request-uri": "http://foo.bar.com/", "request-body": { "cost-mode" : "numerical", "cost-type" : "routingcost", "pids" : { "srcs" : [ "pid1" ], "dsts" : [ "pid1", "pid2", "pid3" ] } }, "media-type": "application/alto-costmap+json", "expires": "20141231" } }, "data": { "map-vtag": "1266506139", "map": { "pid1": { "ipv4": [ "192.0.2.0/24", "198.51.100.0/25" ] }, "pid2": { "ipv4": [ "198.51.100.128/25" ] }, "pid3": { "ipv4": [ "0.0.0.0/0" ], "ipv6": [ "::/0" ] } } } }`)
