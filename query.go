package dql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/influxdata/influxdb1-client/models"
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
func NewClient(dk string) *Client {
	c := &Client{
		dk: dk,
	}

	// TODO: new http client
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
	Series   []*models.Row `json:"series"`
	RawQuery string        `json:"raw_query,omitempty"`
	Cost     string        `json:"cost"`
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
