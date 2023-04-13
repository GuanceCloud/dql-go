// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

import "time"

// DQLOption used to set various DQL options.
type DQLOption func(*dql)

// WithProfile enable profile on query.
func WithProfile(on bool) DQLOption {
	return func(q *dql) {
		q.Profile = on
	}
}

// WithOptimized enable optimize on query.
// For example, if query 7 days data, the query will only response
// the first index data. If order by ASC, the first index is the
// oldest index.
//
// NOTE: not available on M::.
func WithOptimized(on bool) DQLOption {
	return func(q *dql) {
		q.Optimized = on
	}
}

// WithConditions set extra where-condtions to DQL.
func WithConditions(conditions string) DQLOption {
	return func(q *dql) {
		q.Conditions = conditions
	}
}

// WithOutputFormat set output format, currently only support LineProtocol.
func WithOutputFormat(of OutputFormat) DQLOption {
	return func(q *dql) {
		q.OutputFormat = of.String()
	}
}

// WithMaskVisible set visible/hide on sensitive fields.
func WithMaskVisible(on bool) DQLOption {
	return func(q *dql) {
		q.MaskVisible = on
	}
}

// WithAsync set async query.
func WithAsync(on bool) DQLOption {
	return func(q *dql) {
		q.IsAsync = on
	}
}

// WithAsyncTimeout set async query timeout.
func WithAsyncTimeout(du time.Duration) DQLOption {
	return func(q *dql) {
		q.AsyncTimeout = du.String()
	}
}

// WithAsyncID fetch async query result on the id.
func WithAsyncID(id string) DQLOption {
	return func(q *dql) {
		q.AsyncID = id
	}
}

// WithTimeout set query timeout.
func WithTimeout(du time.Duration) DQLOption {
	return func(q *dql) {
		q.Timeout = du.String()
	}
}

// WithShowLabel will show-label in query result.
// NOTE: Only available on query Object(O::).
func WithShowLabel(on bool) DQLOption {
	return func(q *dql) {
		q.ShowLabel = on
	}
}

// WithDisableExpensiveQuery disable/enable expensive query.
// For left wildcard like following will trigger a expensive query.
//
//   L::some_source { f1 = wildcard('*xx') }
//
// NOTE: Disable all expensive query are a good manner to protect
// your worksapce.
func WithDisableExpensiveQuery(on bool) DQLOption {
	return func(q *dql) {
		q.DisableExpensiveQuery = on
	}
}

// WithDisableMultipleField disable/enable query multiple field in single DQL.
func WithDisableMultipleField(on bool) DQLOption {
	return func(q *dql) {
		q.DisableMultipleField = on
	}
}

// WithHighlight enable/disable highlight on query result.
func WithHighlight(on bool) DQLOption {
	return func(q *dql) {
		q.Highlight = on
	}
}

// WithDisableSLimit disable/enable default slimit.
func WithDisableSLimit(on bool) DQLOption {
	return func(q *dql) {
		q.DisableSLimit = on
	}
}

// WithDisableQueryParse disable/enable query parse.
func WithDisableQueryParse(on bool) DQLOption {
	return func(q *dql) {
		q.DisableQueryParse = on
	}
}

// WithMaxDuration set DQL max time range, this option used
// to avoid unexpected too-large query. If time range in DQL
// exceed max duration, there will be a query error returned.
// For example, following DQL will fail if max duration set to
// 1 hour:
//
//   L::some_source [1d:] # query latest 24h logging
//
// We will get a error like:
//
//   parse error: time range should less than 1h0m0s
func WithMaxDuration(du time.Duration) DQLOption {
	return func(q *dql) {
		q.MaxDuration = du.String()
	}
}

// WithMaxPoint used to control max query points under
// group by, for example:
//
//   L::some_source BY status
func WithMaxPoint(n int) DQLOption {
	return func(q *dql) {
		q.MaxPoint = n
	}
}

// WithOffset used to query next page points.
func WithOffset(n int) DQLOption {
	return func(q *dql) {
		q.Offset = n
	}
}

// WithSOffset used to query next page of series.
func WithSOffset(n int) DQLOption {
	return func(q *dql) {
		q.SOffset = n
	}
}

// WithSLimit used to limit max query time series.
func WithSLimit(n int) DQLOption {
	return func(q *dql) {
		q.SLimit = n
	}
}

// WithTimeRange used to set time range of the DQL query.
// start and end are UNIX timestamp in ms.
func WithTimeRange(start, end int) DQLOption {
	return func(q *dql) {
		q.TimeRange = append(q.TimeRange, start)
		q.TimeRange = append(q.TimeRange, end)
	}
}

// WithSearchAfter used to set search-after of the DQL query.
func WithSearchAfter(after ...any) DQLOption {
	return func(q *dql) {
		if len(after) > 0 {
			q.SearchAfter = append(q.SearchAfter, after...)
		}
	}
}

// WithOrderBy used to set order-by of the DQL query.
func WithOrderBy(k string, order OrderByOrder) DQLOption {
	return func(q *dql) {
		q.OrderBy = append(q.OrderBy, OrderBy(map[string]string{k: order.String()}))
	}
}

// QueryOption used to set various query options.
type QueryOption func(*query)

// WithEchoExplain used to echo the translated query of the DQL.
func WithEchoExplain(on bool) QueryOption {
	return func(q *query) {
		q.EchoExplain = on
	}
}

// WithQueries used to send one or more DQLs to a query request.
func WithQueries(arr ...*dql) QueryOption {
	return func(q *query) {
		q.Queries = append(q.Queries, arr...)
	}
}

// WithToken enable we send the query to Dataway directly.
func WithToken(token string) QueryOption {
	return func(q *query) {
		q.Token = token
	}
}

// WithHTTPS are required if we want to send the query
// to public openway.
func WithHTTPS(on bool) QueryOption {
	return func(q *query) {
		q.https = on
	}
}
