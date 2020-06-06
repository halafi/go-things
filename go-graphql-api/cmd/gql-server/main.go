package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/handler"
	"github.com/halafi/go-things/go-graphql-api/pkg/gql"
	"github.com/halafi/go-things/go-graphql-api/pkg/postgres"
	"github.com/halafi/go-things/go-graphql-api/pkg/server"
	"github.com/halafi/go-things/go-graphql-api/pkg/utils"
)

func main() {
	router, db := initializeAPI()
	defer db.Close(context.Background())
	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	db, err := postgres.New(
		utils.MustGet("DB_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	sc, err := gql.CreateSchema(db)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	r := chi.NewRouter()
	// Add some middleware to our router
	r.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,       // log api request calls
		middleware.Compress(5),  // compress results, mostly gzipping assets and json
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)
	h := handler.New(&handler.Config{
		Schema:   &sc,
		Pretty:   true,
		GraphiQL: true,
	})
	// Create the graphql route with a Server method to handle it
	r.Post("/graphql", s.GraphQL())
	r.Mount("/graphiql", h)

	return r, db
}
