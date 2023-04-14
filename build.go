// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

// MustBuildDQL build a DQL query, and panic on any error.
// Most of the time, the build will not fail, we can use
// the Must function without worry.
func MustBuildDQL(dql string, opts ...DQLOption) *dql {
	q, err := BuildDQL(dql, opts...)
	if err != nil {
		panic(err.Error())
	}
	return q
}

// BuildDQL used to build a DQL query with one or more options.
// dqlStr is the basic DQL query string.
func BuildDQL(dqlStr string, opts ...DQLOption) (*dql, error) {
	q := &dql{
		DQL:         dqlStr,
		SearchAfter: []any{}, // default enable search-after
	}

	for _, opt := range opts {
		if opt != nil {
			opt(q)
		}
	}

	return q, nil
}
