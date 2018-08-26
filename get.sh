#!/bin/sh

# This is a convenience shell script to load packages
# you'll likely need with a GraphQL/Dgraph project

go get -u github.com/brettallred/go-logit
go get -u	github.com/dgraph-io/dgo
go get -u github.com/dgraph-io/dgo/protos/api

go get -u github.com/graphql-go/graphql
go get -u github.com/graphql-go/graphql/testutil
go get -u github.com/graphql-go/handler

go get -u github.com/stretchr/testify/assert
go get -u google.golang.org/grpc
