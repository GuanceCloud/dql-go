// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

import "time"

// DQLOption used to set various DQL options.
type DQLOption func(*dql)

// WithProfile enable profiling. Only available for OpenSearch backend.
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
//
// Deprecated: this option is dropped.
func WithShowLabel(on bool) DQLOption {
	return func(q *dql) {
		q.ShowLabelDeprecated = on
	}
}

// WithDisableExpensiveQuery disable/enable expensive query.
// For left wildcard like following will trigger a expensive query.
//
//	L::some_source { f1 = wildcard('*xx') }
//
// NOTE: Disable all expensive query are a good manner to protect
// your worksapce.
//
// Deprecated: Option removed for DQL.
func WithDisableExpensiveQuery(on bool) DQLOption {
	return func(q *dql) {
		q.DisableExpensiveQueryDeprecated = on
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
//	L::some_source [1d:] # query latest 24h logging
//
// We will get a error like:
//
//	parse error: time range should less than 1h0m0s
func WithMaxDuration(du time.Duration) DQLOption {
	return func(q *dql) {
		q.MaxDuration = du.String()
	}
}

// WithMaxPoint used to control max query points under
// group by, for example:
//
//	L::re(`.*`):(fill(count(__docid), 0) AS count) [1d] BY status
//
// All logs are group by its status, each status(bucket) may have different
// number of logs, and we can limit only n points in each status.
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

// WithOrderBy used to set order-by on point.
func WithOrderBy(k string, order OrderByOrder) DQLOption {
	return func(q *dql) {
		q.OrderBy = append(q.OrderBy, orderBy(map[string]string{k: order.String()}))
		q.NOrderBy = append(q.NOrderBy, orderBy(map[string]string{k: order.String()}))
	}
}

// WithSOrderBy used to set order-by on series.
func WithSOrderBy(k string, order OrderByOrder) DQLOption {
	return func(q *dql) {
		q.NSOrderBy = append(q.NSOrderBy, orderBy(map[string]string{k: order.String()}))
	}
}

// WithSampling used to enable/disable sampling of the query result.
func WithSampling(on bool) DQLOption {
	return func(q *dql) {
		q.DisableSampling = !on
	}
}

// WithQueryType set query type, only "dql" or "promql" are allowed.
func WithQueryType(t string) DQLOption {
	return func(q *dql) {
		switch t {
		case "dql", "promql":
			q.QType = t
		default: // pass
		}
	}
}

// WithAlignTime enable time alignment for query result.
func WithAlignTime(on bool) DQLOption {
	return func(q *dql) {
		q.AlignTime = on
	}
}

// WithLargeQuery enable/disable large-data-set query. Large-data-set query default
// enabled, but we can disable it to protect backend storage.
func WithLargeQuery(on bool) DQLOption {
	return func(q *dql) {
		q.DisallowLargeQuery = !on
	}
}

// WithCursorTime set cursort timestamp for paging. The timestamp n can be s/ms/us,
// and the backend will guess the unit of the timestamp.
func WithCursorTime(n int64) DQLOption {
	return func(q *dql) {
		if n > 0 {
			q.CursorTime = n
		}
	}
}

// WithStepInterval set time step interval for time-aggregate query.
func WithStepInterval(n int64) DQLOption {
	return func(q *dql) {
		if n > 0 {
			q.Interval = n
		}
	}
}

// WithLimit set max returned point's. Default is 1000.
func WithLimit(n int64) DQLOption {
	return func(q *dql) {
		if n > 0 {
			q.Limit = n
		}
	}
}

// WithRoleRules set role rules for the query.
//
// Rules example:
//
//	rules := map[string][]QueryRule{
//		"logging": []QueryRule{ // rule for query logging(aka L::) data.
//		QueryRule{
//				Rule:  `fruit IN ['apple', 'orange']`,
//				Index: []string{"my-logging-index-name"},
//			},
//		},
//	}
func WithRoleRules(rules map[string][]QueryRule) DQLOption {
	return func(q *dql) {
		q.Rules = rules
	}
}

func WithMultipleIndices(idx ...*DorisIndices) DQLOption {
	return func(q *dql) {
		q.Indices = append(q.Indices, idx...)
	}
}

// WithMultipleWorkspaceRules query among multiple workspaces.
func WithMultipleWorkspaceRules(rules ...*WorkspaceIndexRule) DQLOption {
	return func(q *dql) {
		q.IndexList = append(q.IndexList, rules...)
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
