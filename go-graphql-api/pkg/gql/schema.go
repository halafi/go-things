package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/halafi/go-things/go-graphql-api/pkg/postgres"
)

// CreateSchema creates grafql schema
func CreateSchema(db *postgres.Db) (graphql.Schema, error) {
	// Create our root query for graphql
	rootQuery := NewRoot(db)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	return sc, err
}
