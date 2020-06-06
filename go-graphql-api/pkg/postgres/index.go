package postgres

import (

	// postgres driver
	"context"

	"github.com/jackc/pgx/v4"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*pgx.Conn
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	// db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	return &Db{conn}, nil
}
