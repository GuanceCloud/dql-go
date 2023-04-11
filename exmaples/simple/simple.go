package main

import (
	"encoding/json"
	"log"

	"github.com/GuanceCloud/dql-go"
)

func main() {

	c := dql.NewClient("localhost:9529")

	r, err := c.Query(dql.WithQueries(dql.MustBuildDQL("M::cpu LIMIT 1")))

	if err != nil {
		panic(err.Error())
	}

	j, err := json.MarshalIndent(r, "", "  ")
	log.Printf("result: %s", string(j))
}
