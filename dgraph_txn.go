package agora

import (
	"context"
	"flag"
	"log"

	logit "github.com/brettallred/go-logit"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var (
	dgraphHost = flag.String("d", "127.0.0.1:9080", "Dgraph server address")
)

// Dial is a helper to create a DGraph connection
func Dial() *grpc.ClientConn {
	conn, err := grpc.Dial(*dgraphHost, grpc.WithInsecure())
	if err != nil {
		logit.Fatal(" While trying to dial gRPC")
	}

	return conn
}

// MutateDgraph is a helper func to run Mutate operations on Dgraph
func MutateDgraph(conn *grpc.ClientConn, jsonStr []byte) *api.Assigned {
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	ctx := context.Background()
	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   jsonStr,
	}

	assigned, err := txn.Mutate(ctx, mu)
	if err != nil {
		logit.Fatal(" %e", err)
	}
	return assigned
}

// QueryDgraph runs a query operation on DGraph and returns the JSON response
func QueryDgraph(conn *grpc.ClientConn, query string) []byte {
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	ctx := context.Background()
	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	resp, err := txn.Query(ctx, query)
	if err != nil {
		logit.Fatal(" Query Error: %e", err)
	}

	return resp.Json
}

// AlterDgraph runs an alter Dgraph operation
func AlterDgraph(conn *grpc.ClientConn, op *api.Operation) {
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()
	err := dg.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}
}
