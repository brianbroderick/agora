package agora

import (
	"context"
	"os"

	"github.com/brianbroderick/logit"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

func GetDgraphHost() string {
	dh := os.Getenv("DGRAPH_HOST")
	if dh == "" {
		return "127.0.0.1:9080"
	}
	return dh
}

// NewDgraphConn establishes a new Dgraph Connection
func NewDgraphConn() *DgraphConn {
	c := &DgraphConn{}
	c.EstablishConn()
	return c
}

// NewDgraphTxn establishes a new Dgraph Connection and Transaction
func NewDgraphTxn() *DgraphConn {
	c := NewDgraphConn()
	c.EstablishTxn()
	return c
}

// DgraphConn is a struct holding connection data
type DgraphConn struct {
	Conn *grpc.ClientConn
	Dc   api.DgraphClient
	Dg   *dgo.Dgraph
	Ctx  context.Context
	Txn  *dgo.Txn
	Err  error
}

// EstablishConn establishes a Dgraph connection
func (c *DgraphConn) EstablishConn() {
	c.Conn = Dial()
}

// EstablishTxn establishes a Dgraph transaction
func (c *DgraphConn) EstablishTxn() {
	c.Dc = api.NewDgraphClient(c.Conn)
	c.Dg = dgo.NewDgraphClient(c.Dc)
	c.Ctx = context.Background()
	c.Txn = c.Dg.NewTxn()
}

// DiscardTxn discards a Dgraph transaction and closes the connection
func (c *DgraphConn) DiscardTxn() {
	c.Txn.Discard(c.Ctx)
	c.Conn.Close()
}

// DiscardConn discards a Dgraph connection
func (c *DgraphConn) DiscardConn() {
	c.Conn.Close()
}

// Dial is a helper to create a DGraph connection
func Dial() *grpc.ClientConn {
	dh := GetDgraphHost()
	conn, err := grpc.Dial(dh, grpc.WithInsecure())
	if err != nil {
		logit.Fatal(" While trying to dial gRPC")
	}

	return conn
}
