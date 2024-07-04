// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/GuanceCloud/dql-go"
)

func main() {
	c := dql.NewClient("localhost:9529")

	r, err := c.Query(dql.WithQueries(
		dql.MustBuildDQL("M::cpu LIMIT 1",
			dql.WithTimeout(time.Second*30))))
	if err != nil {
		panic(err.Error())
	}

	j, _ := json.MarshalIndent(r, "", "  ")
	log.Printf("result: %s", string(j))
}
