package agora

import (
	"log"

	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// DropAll drops the entire Dgraph database
func DropAll() {
	conn, err := grpc.Dial(*dgraphHost, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	AlterDgraph(conn, &api.Operation{DropAll: true})
}

// SetSchema accepts a schema string and sets it in Dgraph
func SetSchema(schema string) {
	conn, err := grpc.Dial(*dgraphHost, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	op := &api.Operation{}
	op.Schema = schema

	AlterDgraph(conn, op)
}
