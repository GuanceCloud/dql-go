// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package dql wraps DQL query SDK.
package dql

type dql struct {
	SearchAfter  []any  `json:"search_after"`
	TimeRange    []int  `json:"time_range,omitempty"`
	Conditions   string `json:"conditions,omitempty"`
	DQL          string `json:"query"`
	MaxDuration  string `json:"max_duration,omitempty"`
	AsyncTimeout string `json:"async_timeout,omitempty"`
	Timeout      string `json:"search_timeout,omitempty"`
	OutputFormat string `json:"output_format,omitempty"`
	AsyncID      string `json:"async_id"`

	OrderBy []OrderBy `json:"orderby,omitempty"`

	SLimit   int `json:"slimit,omitempty"`
	SOffset  int `json:"soffset,omitempty"`
	Offset   int `json:"offset,omitempty"`
	MaxPoint int `json:"max_point,omitempty"`

	DisableQueryParse     bool `json:"disable_query_parse,omitempty"`
	DisableSLimit         bool `json:"disable_slimit,omitempty"`
	Highlight             bool `json:"highlight,omitempty"`
	DisableMultipleField  bool `json:"disable_multiple_field,omitempty"`
	DisableExpensiveQuery bool `json:"disable_expensive_query,omitempty"`
	ShowLabel             bool `json:"show_label,omitempty"`
	IsAsync               bool `json:"is_async,omitempty"`
	MaskVisible           bool `json:"mask_visible,omitempty"`
	Profile               bool `json:"is_profile,omitempty"`
	Optimized             bool `json:"is_optimized,omitempty"`
	DisableSampling       bool `json:"disable_sampling,omitempty"`

	// TODO: following args not support
	//AlignTime            bool                   `json:"align_time"`  // guancedb自动对齐
	//CursorTime           int64                  `json:"cursor_time"` // doris分段查询阀值
	//Date                 string                 `json:"date"`
	//DisallowLargeQuery   bool                   `json:"disallow_large_query"`   // doris 使用耗时查询
	//EnableExpensiveQuery bool                   `json:"enable_expensive_query"` // use expensive query, default false, can not use
	//IndexList            []*WorkspaceIndexRule  `json:"index_list,omitempty"`   // 跨空间查询条件
	//Indices              []*DorisIndices        `json:"indices,omitempty"`      // 传递给doris的跨空间查询条件
	//Interval             int64                  `json:"interval"`
	//IsOptimized          bool                   `json:"is_optimized"` // search really indicies
	//IsProfile            bool                   `json:"is_profile"`   // profile search
	//Label                bool                   `json:"show_label"`   // 是否展示labels
	//Limit                int64                  `json:"limit"`
	//MaxScanSize          int64                  `json:"max_scan_size"` // scan completed
	//NOrderBy             []map[string]string    `json:"order_by"`
	//NSOrderBy            []map[string]string    `json:"sorder_by"`
	//QType                string                 `json:"qtype"` // dql or promql
	//Rules                map[string][]QueryRule `json:"rules"` // rules
	//SOrderBy             []map[string]string    `json:"sorderby"`
	//ScanCompleted        bool                   `json:"scan_completed"` // scan completed
	//ScanIndex            string                 `json:"scan_index"`     // scan index
	//SearchTimeout        string                 `json:"search_timeout"` // 最大执行时间
	//Timezone             string                 `json:"tz"`
}
