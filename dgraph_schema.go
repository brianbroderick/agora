package agora

import (
	"github.com/dgraph-io/dgo/v200/protos/api"
)

// DropAll drops the entire Dgraph database
func DropAll() {
	AlterDgraph(&api.Operation{DropAll: true})
}

// SetSchema accepts a schema string and sets it in Dgraph
func SetSchema(schema string) {
	op := &api.Operation{}
	op.Schema = schema

	AlterDgraph(op)
}
