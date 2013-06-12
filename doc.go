// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

// Package alto implements JSON encoder and decoder for the
// Application-Layer Traffic Optimization (ALTO) protocol as described
// in http://tools.ietf.org/html/draft-ietf-alto-protocol.
//
//
// Encoding at server side:
//
//	accepts := req.Header.Get("Accept")
//	if accepts == "" {
//		// error handling
//	}
//	if !strings.Contains(accepts, alto.MediaTypeDirectory) {
//		// error handling
//	}
//	w.Header().Set("Content-Type", alto.MediaTypeDirectory)
//	var dir alto.Directory
//	if err := json.NewEncoder(w).Encode(&dir); err != nil {
//		// error handling
//	}
//
//
// Decoding at client side:
//
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", "http://...", nil)
//	if err != nil {
//		// error handling
//	}
//	req.Header.Set("Accept", alto.MediaTypeNetworkMap+","+alto.MediaTypeError)
//	resp, err := client.Do(req)
//	if err != nil {
//		// error handling
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode != http.StatusOK {
//		// error handling
//	}
//	nmap := alto.NewResource("networkmap")
//	if err := json.NewDecoder(resp.Body).Decode(nmap); err != nil {
//		// error handling
//	}
package alto
