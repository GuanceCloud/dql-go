// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package dql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type query struct {
	https bool

	EchoExplain bool   `json:"echo_explain"`
	Token       string `json:"token,omitempty"`
	Queries     []*dql `json:"queries"`
}

func (q *query) json(indent bool) string {
	if indent {
		j, err := json.MarshalIndent(q, "", " ")
		if err != nil {
			return ""
		}
		return string(j)
	}

	j, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(j)
}

// A Client is the DQL query client connecting to a exist Datakit
// or directly to Dataway(and the token required).
type Client struct {
	host   string
	cli    *http.Client
	dqlURL string

	lastQuery *query
}

// NewClient create a Datakit/Dataway client with IP:Port.
// For example, local default Datakit host is localhost:9529, for
// directly to dataway, the default host is openway.guance.com.
func NewClient(host string) *Client {
	c := &Client{
		host: host,
	}

	c.cli = &http.Client{}

	return c
}

// A Result is the query result of DQL request. Within a Result
// there maybe multiple DQL query result.
// If there any error on query, we can see them with ErrorCode and Message.
type Result struct {
	ErrorCode string       `json:"error_code,omitempty"`
	Message   string       `json:"message,omitempty"`
	Content   []*DQLResult `json:"content"`
}

type AsyncSearchTaskPayload struct {
	CreateTime    time.Time
	SearchTimeout string
	AsyncID       string
	Timeout       string
	Wsuuid        string
}

// A DQLResult is a single DQL's query result.
type DQLResult struct {
	Series []*Row `json:"series"`

	// Base64 encoded lineprotocol output
	Points []string `json:"points"`

	GroupByList []string `json:"group_by,omitempty"`

	SearchAfter []interface{} `json:"search_after,omitempty"`

	Cost         string      `json:"cost"`
	RawQuery     string      `json:"raw_query,omitempty"`
	QueryParse   interface{} `json:"query_parse,omitempty"`
	QueryWarning string      `json:"query_warning,omitempty"`

	Totalhits   int64 `json:"total_hits,omitempty"`
	FilterCount int64 `json:"filter_count,omitempty"`

	// Async query ID
	AsyncID string `json:"async_id,omitempty"`

	// Logging index name
	IndexName string `json:"index_name"`

	// Logging storage type(sls/outer_sls/es)
	IndexStoreType string `json:"index_store_type"`

	// Query type(influxdb/tdengine/guancedb)
	QueryType string `json:"query_type"`

	// Async query still running or not
	IsRunning  bool   `json:"is_running"`
	Complete   bool   `json:"complete"`
	IndexNames string `json:"index_names"` // index names
}

// Row represents a single row returned from the execution of a statement.
type Row struct {
	Name    string            `json:"name,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
	Columns []string          `json:"columns,omitempty"`
	Values  [][]interface{}   `json:"values,omitempty"`
	Partial bool              `json:"partial,omitempty"`
}

// Query send one or more DQL query to Datakit. We can build
// DQL within QueryOptions.
func (c *Client) Query(opts ...QueryOption) (*Result, error) {
	q := &query{}

	for _, opt := range opts {
		if opt != nil {
			opt(q)
		}
	}

	return c.do(q)
}

func (c *Client) do(q *query) (*Result, error) {
	j, err := json.Marshal(q)
	if err != nil {
		c.lastQuery = nil
		return nil, err
	}

	c.lastQuery = q

	if c.dqlURL == "" {
		if q.https {
			c.dqlURL = fmt.Sprintf("https://%s/v1/query/raw", c.host)
		} else {
			c.dqlURL = fmt.Sprintf("http://%s/v1/query/raw", c.host)
		}

		if q.Token != "" {
			c.dqlURL = fmt.Sprintf("%s?token=%s", c.dqlURL, q.Token)
		}
	}

	req, err := http.NewRequest("POST", c.dqlURL, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //nolint:errcheck

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r Result
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
