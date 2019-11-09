package agora

import (
	"github.com/brianbroderick/dgo/v2/protos/api"
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
