// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

import (
	T "testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *T.T) {
	t.Run(`with-x`, func(t *T.T) {
		q := MustBuildDQL("L::nginx", WithProfile(true))
		assert.True(t, q.Profile)

		q = MustBuildDQL("L::nginx", WithOptimized(true))
		assert.True(t, q.Optimized)

		q = MustBuildDQL("L::nginx", WithConditions("some-condition"))
		assert.Equal(t, "some-condition", q.Conditions)

		q = MustBuildDQL("L::nginx", WithOutputFormat(LineProtocol))
		assert.Equal(t, LineProtocol.String(), q.OutputFormat)
	})
}
