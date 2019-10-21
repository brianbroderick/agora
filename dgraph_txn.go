package agora

import (
	"encoding/json"
	"log"

	logit "github.com/brettallred/go-logit"
	"github.com/dgraph-io/dgo/protos/api"
)

// MutateDgraph is a helper func to run Mutate operations on Dgraph
func MutateDgraph(jsonStr []byte) *api.Response {
	c := NewDgraphTxn()
	defer c.DiscardTxn()

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   jsonStr,
	}

	assigned, err := c.Txn.Mutate(c.Ctx, mu)
	if err != nil {
		logit.Fatal(" %e", err)
	}
	return assigned
}

// QueryDgraph runs a query operation on DGraph and returns the JSON response
func QueryDgraph(query string) []byte {
	c := NewDgraphTxn()
	defer c.DiscardTxn()

	resp, err := c.Txn.Query(c.Ctx, query)
	if err != nil {
		logit.Fatal(" Query Error: %e", err)
	}

	return resp.Json
}

// QueryDgraphWithVars runs a query operation on DGraph with variables and returns the JSON response
func QueryDgraphWithVars(query string, variables map[string]string) []byte {
	c := NewDgraphTxn()
	defer c.DiscardTxn()

	resp, err := c.Txn.QueryWithVars(c.Ctx, query, variables)
	if err != nil {
		logit.Fatal(" Query Error: %e", err)
	}

	return resp.Json
}

func ResolveQuery(r interface{}, query string) error {
	j := QueryDgraph(query)

	err := json.Unmarshal(j, &r)
	if err != nil {
		return err
	}

	return nil
}

func ResolveQueryWithVars(r interface{}, query string, variables map[string]string) error {
	j := QueryDgraphWithVars(query, variables)

	err := json.Unmarshal(j, &r)
	if err != nil {
		return err
	}

	return nil
}

// AlterDgraph runs an alter Dgraph operation
func AlterDgraph(op *api.Operation) {
	c := NewDgraphTxn()
	defer c.DiscardTxn()

	err := c.Dg.Alter(c.Ctx, op)
	if err != nil {
		log.Fatal(err)
	}
}
