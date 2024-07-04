// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package dql wraps DQL query SDK.
package dql

type QueryRule struct {
	Rule  string   `json:"rule"`
	Index []string `json:"indexes"`
}

type WorkspaceIndexRule struct {
	WorkspaceUUID string                 `json:"workspace_uuid"`
	IndexName     string                 `json:"index_name"`
	Rules         map[string][]QueryRule `json:"rules"`
}

type DorisIndices struct {
	TenantID  string `json:"tenant_id"`
	IndexName string `json:"index_name"`
	Condition string `json:"conditions"`
}

// DQL query options defined in(internal only)
//
//	https://confluence.jiagouyun.com/pages/viewpage.action?pageId=196018193
type dql struct {
	SearchAfter []any                 `json:"search_after"`
	TimeRange   []int                 `json:"time_range,omitempty"`
	OrderBy     []orderBy             `json:"orderby,omitempty"`
	NOrderBy    []map[string]string   `json:"order_by"`  // the newer order-by
	NSOrderBy   []map[string]string   `json:"sorder_by"` // the newer sorder-by
	IndexList   []*WorkspaceIndexRule `json:"index_list,omitempty"`
	Indices     []*DorisIndices       `json:"indices,omitempty"`

	Conditions   string `json:"conditions,omitempty"`
	DQL          string `json:"query"`
	MaxDuration  string `json:"max_duration,omitempty"`
	AsyncTimeout string `json:"async_timeout,omitempty"`
	Timeout      string `json:"search_timeout,omitempty"`
	OutputFormat string `json:"output_format,omitempty"`
	AsyncID      string `json:"async_id"`
	QType        string `json:"qtype"` // dql or promql

	SLimit     int                    `json:"slimit,omitempty"`
	SOffset    int                    `json:"soffset,omitempty"`
	Offset     int                    `json:"offset,omitempty"`
	MaxPoint   int                    `json:"max_point,omitempty"`
	CursorTime int64                  `json:"cursor_time"`
	Interval   int64                  `json:"interval"`
	Limit      int64                  `json:"limit"`
	Rules      map[string][]QueryRule `json:"rules"` // rules

	DisableQueryParse    bool `json:"disable_query_parse,omitempty"`
	DisableSLimit        bool `json:"disable_slimit,omitempty"`
	Highlight            bool `json:"highlight,omitempty"`
	DisableMultipleField bool `json:"disable_multiple_field,omitempty"`

	DisableExpensiveQueryDeprecated bool `json:"disable_expensive_query,omitempty"`
	ShowLabelDeprecated             bool `json:"show_label,omitempty"`

	IsAsync            bool `json:"is_async,omitempty"`
	MaskVisible        bool `json:"mask_visible,omitempty"`
	Profile            bool `json:"is_profile,omitempty"`
	Optimized          bool `json:"is_optimized,omitempty"`
	DisableSampling    bool `json:"disable_sampling,omitempty"`
	AlignTime          bool `json:"align_time"`
	DisallowLargeQuery bool `json:"disallow_large_query"`
}
