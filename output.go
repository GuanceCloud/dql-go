// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

// A OutputFormat is the query result format.
type OutputFormat int

// Query result formats.
const (
	LineProtocol OutputFormat = iota
)

// String used to get the output format in string representation.
func (of OutputFormat) String() string {
	switch of {
	case LineProtocol:
		return "lineprotocol"
	}
	return ""
}
