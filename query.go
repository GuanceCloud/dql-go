package dql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type query struct {
	Queries     []*dql `json:"queries"`
	EchoExplain bool   `json:"echo_explain"`
}

// A Client is the DQL query client connecting to a exist Datakit.
type Client struct {
	dk  string
	cli *http.Client
}

// NewClient create a datakit client with IP:Port.
// For example, local default Datakit host is localhost:9529.
func NewClient(dk string) *Client {
	c := &Client{
		dk: dk,
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

// A DQLResult is a single DQL's query result.
type DQLResult struct {
	Series   []*Row `json:"series"`
	RawQuery string `json:"raw_query,omitempty"`
	Cost     string `json:"cost"`
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
	j, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://%s/v1/query/raw", c.dk),
		bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r Result
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
