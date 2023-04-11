// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

// A OrderByOrder is the order-by option, DESC or ASC
type OrderByOrder int

// String() is the string-representation of DESC and ASC
func (o OrderByOrder) String() string {
	switch o {
	case ASC:
		return "asc"
	case DESC:
		return "desc"
	}
	return ""
}

// The order-by options: asc and desc
const (
	ASC OrderByOrder = iota
	DESC
)
