// Copyright 2013 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.

package alto

import "errors"

var (
	errUnknownAddress = errors.New("unknown address")
)

const (
	MediaTypeError = "application/alto-error+json" // media type for ALTO error notification
)

const (
	ErrSyntax              = "E_SYNTAX"
	ErrJSONFieldMissing    = "E_JSON_FIELD_MISSING"
	ErrJSONValueType       = "E_JSON_VALUE_TYPE"
	ErrInvalidCostMode     = "E_INVALID_COST_MODE"
	ErrInvalidCostMetric   = "E_INVALID_COST_METRCI"
	ErrInvalidPropertyType = "E_INVALID_PROPERTY_TYPE"
)

// An Error represents an error notification.
type Error struct {
	Code string `json:"code"`
}
