// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

import (
	"encoding/base64"
	"encoding/json"
	"os"
	T "testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryDataway(t *T.T) {
	// run these test with:
	//   DK_TOKEN="tkn_xxxxxx" go test -test.v -run TestQueryDataway
	token := os.Getenv("DK_TOKEN")
	if token == "" {
		t.Skipf("%s skipped due to DK_TOKEN not set", t.Name())
	}

	c := NewClient("openway.guance.com")

	t.Run("with-max-point", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::re(`.*`):(fill(count(__docid), 0) AS count) [1d] BY status",
					WithMaxPoint(2))))

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)
		t.Logf("resp:\n%s", string(j))

		assert.NoError(t, err)

		t.Logf("request:\n%s\n---\nseries: %d, cost: %s, client-cost: %s",
			c.lastQuery.json(true),
			len(r.Content[0].Series),
			r.Content[0].Cost,
			time.Since(start),
		)

		for _, s := range r.Content[0].Series {
			t.Logf("pts: %d", len(s.Values))
		}
	})

	t.Run("with-max-duration", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module:(f1,f2) [30d:]",
					WithMaxDuration(time.Hour),
				)))

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)
		t.Logf("resp:\n%s", string(j))

		assert.NoError(t, err)
		assert.NotEmpty(t, r.ErrorCode)
		assert.Empty(t, r.Content)

		t.Logf("request:\n%s\n---\n client-cost: %s",
			c.lastQuery.json(true),
			time.Since(start),
		)
	})

	t.Run("disable-multi-field", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module:(f1,f2) limit 1",
					WithDisableMultipleField(true),
				)))

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)
		t.Logf("resp:\n%s", string(j))

		assert.NoError(t, err)
		assert.NotEmpty(t, r.ErrorCode)
		assert.Empty(t, r.Content)

		t.Logf("request:\n%s\n---\n client-cost: %s",
			c.lastQuery.json(true),
			time.Since(start),
		)
	})

	t.Run("with-time-range", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module",
					WithTimeRange(1680172008117, 1680172108117), // timestamp ms
				)))

		assert.NoError(t, err)
		require.NotEmpty(t, r.Content[0].Series)

		t.Logf("request:\n%s\n---\nresult: %d, cost: %s, client-cost: %s",
			c.lastQuery.json(true),
			len(r.Content[0].Series[0].Values),
			r.Content[0].Cost,
			time.Since(start),
		)
	})

	t.Run("disalbe-expensive-query", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module LIMIT 10000",
					WithDisableExpensiveQuery(true),
				)))

		assert.NoError(t, err)
		require.NotEmpty(t, r.Content[0].Series)

		t.Logf("request:\n%s\n---\nresult: %d, cost: %s, client-cost: %s",
			c.lastQuery.json(true),
			len(r.Content[0].Series[0].Values),
			r.Content[0].Cost,
			time.Since(start),
		)
	})

	t.Run("order-by", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module:(name, coverage, message) { coverage > 0 } LIMIT 3",
					WithOrderBy("coverage", ASC),
				)))

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))

		require.NotEmpty(t, r.Content[0].Series)

		t.Logf("request:\n%s\n---\nresult: %d, cost: %s, client-cost: %s",
			c.lastQuery.json(true),
			len(r.Content[0].Series[0].Values),
			r.Content[0].Cost,
			time.Since(start),
		)
	})

	t.Run("single", func(t *T.T) {
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("M::cpu limit 1"),
			))

		assert.NoError(t, err)
		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("logging", func(t *T.T) {
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module limit 1"),
			))

		assert.NoError(t, err)
		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("logging-searching-aster", func(t *T.T) {
		sa := []any{}
		for i := 0; i < 3; i++ {
			r, err := c.Query(
				WithEchoExplain(true),
				WithHTTPS(true),
				WithToken(token),
				WithQueries(
					MustBuildDQL("L::testing_module limit 1",
						WithSearchAfter(sa...),
					),
				))
			assert.NoError(t, err)

			if len(r.Content) > 0 {
				sa = r.Content[0].SearchAfter

				j, err := json.MarshalIndent(r, "", "  ")
				assert.NoError(t, err)

				t.Logf("result:\n%s", string(j))
			}
		}
	})

	t.Run("logging-profile", func(t *T.T) {
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module limit 1", WithProfile(true)),
			))
		assert.NoError(t, err)

		assert.NotEmpty(t, r.Content[0].IndexNames)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("logging-optimized", func(t *T.T) {
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module limit 1", WithOptimized(true)),
			))
		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("logging-condition", func(t *T.T) {
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module limit 1",
					WithConditions("cost > 10000000000 and status='pass'"),
				),
			))
		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("logging-output-lineprotocol", func(t *T.T) {
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module limit 1",
					WithConditions("cost > 10000000000 and status='pass'"),
					WithOutputFormat(LineProtocol),
				),
			))
		assert.NoError(t, err)

		if len(r.Content) > 0 {
			pts := r.Content[0].Points
			assert.NotEmpty(t, pts)

			for _, pt := range pts {
				decoded, err := base64.StdEncoding.DecodeString(pt)
				assert.NoError(t, err)
				t.Logf("%s", string(decoded))
			}
		}

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("logging-with-timeout", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::http_dial_testing limit 10000",
					WithTimeout(time.Second*3), // fast timeout
				),
			))

		assert.NoError(t, err)

		if r.ErrorCode != "" {
			j, err := json.MarshalIndent(r, "", "  ")
			assert.NoError(t, err)

			t.Logf("result:\n%s", string(j))
			return
		}

		assert.NotEmpty(t, r.Content)
		require.NotEmpty(t, r.Content[0].Series)

		t.Logf("request:\n%s\n---\nresult: %d, cost: %s, client-cost: %s",
			c.lastQuery.json(true),
			len(r.Content[0].Series[0].Values),
			r.Content[0].Cost,
			time.Since(start),
		)
	})

	t.Run("logging-async", func(t *T.T) {
		start := time.Now()
		r, err := c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module limit 10000", WithAsync(true)),
			))
		assert.NoError(t, err)

		asyncID := r.Content[0].AsyncID
		assert.NotEmpty(t, asyncID)
		require.NotEmpty(t, r.Content[0].Series)

		t.Logf("request:\n%s\n---\nresult: %d, cost: %s, client-cost: %s",
			c.lastQuery.json(true),
			len(r.Content[0].Series[0].Values),
			r.Content[0].Cost,
			time.Since(start),
		)

		start = time.Now()
		r, err = c.Query(
			WithEchoExplain(true),
			WithHTTPS(true),
			WithToken(token),
			WithQueries(
				MustBuildDQL("L::testing_module limit 10000", WithAsyncID(asyncID)),
			))
		assert.NoError(t, err)
		require.NotEmpty(t, r.Content[0].Series)

		t.Logf("request:\n%s\n---\nresult: %d, cost: %s, client-cost: %s",
			c.lastQuery.json(true),
			len(r.Content[0].Series[0].Values),
			r.Content[0].Cost,
			time.Since(start),
		)
	})
}

func TestQuery(t *T.T) {
	t.Run("single-dql", func(t *T.T) {
		c := NewClient("localhost:9529")

		r, err := c.Query(WithEchoExplain(true),
			WithQueries(
				MustBuildDQL("M::cpu LIMIT 1"),
			),
		)

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("multi-dql", func(t *T.T) {
		c := NewClient("localhost:9529")

		r, err := c.Query(WithEchoExplain(true),
			WithQueries(
				MustBuildDQL("M::cpu LIMIT 1"),
				MustBuildDQL("M::mem LIMIT 1"),
			),
		)

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("dql-with-option", func(t *T.T) {
		c := NewClient("localhost:9529")

		r, err := c.Query(WithEchoExplain(true),
			WithQueries(
				MustBuildDQL("L::testing_module LIMIT 1"),
				MustBuildDQL("L::testing_module LIMIT 1", WithOffset(10)),
			),
		)

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})

	t.Run("dql-with-page-rotate", func(t *T.T) {
		c := NewClient("localhost:9529")

		pages := 10

		for i := 0; i < pages; i++ {
			r, err := c.Query(WithEchoExplain(true),
				WithQueries(
					MustBuildDQL("L::testing_module LIMIT 2", WithOffset(i*2)),
				),
			)

			assert.NoError(t, err)
			require.NotEmpty(t, r.Content[0].Series)
			t.Logf("get %d", len(r.Content[0].Series[0].Values))
		}
	})

	t.Run("highlight", func(t *T.T) {
		c := NewClient("localhost:9529")

		r, err := c.Query(WithEchoExplain(true),
			WithQueries(
				MustBuildDQL("L::testing_module { message=query_string('datakit') } LIMIT 2", WithHighlight(true)),
			),
		)

		assert.NoError(t, err)

		j, err := json.MarshalIndent(r, "", "  ")
		assert.NoError(t, err)

		t.Logf("result:\n%s", string(j))
	})
}
