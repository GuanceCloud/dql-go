// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

type dql struct {
	Conditions string `json:"conditions,omitempty"`
	DQL        string `json:"query"`

	SearchAfter []any `json:"search_after,omitempty"`

	OrderBy map[string]string `json:"orderby,omitempty"`

	TimeRange []int `json:"time_range,omitempty"`

	SLimit   int `json:"slimit,omitempty"`
	SOffset  int `json:"soffset,omitempty"`
	Offset   int `json:"offset,omitempty"`
	MaxPoint int `json:"max_point,omitempty"`

	MaxDuration string `json:"max_duration,omitempty"`

	DisableQueryParse     bool `json:"disable_query_parse,omitempty"`
	DisableSLimit         bool `json:"disable_slimit,omitempty"`
	Highlight             bool `json:"highlight,omitempty"`
	DisableMultipleField  bool `json:"disable_multiple_field,omitempty"`
	DisableExpensiveQuery bool `json:"disable_expensive_query,omitempty"`
	ShowLabel             bool `json:"show_label,omitempty"`

	IsAsync      bool   `json:"is_async,omitempty"`
	AsyncTimeout string `json:"async_timeout,omitempty"`

	MaskVisible  bool   `json:"mask_visible,omitempty"`
	OutputFormat string `json:"output_format,omitempty"`
}
