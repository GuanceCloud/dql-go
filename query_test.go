package dql

import (
	"encoding/json"
	T "testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *T.T) {
	//t.Skip() // we should test under localhost datakit ok.

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
