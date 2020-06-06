package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

// User shape
type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func processRows(rows pgx.Rows) []User {
	var r User
	// Create slice of Users for our response
	users := []User{}
	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() {
		err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			log.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}
	return users
}

// GetUsers retrieves all users
func (d *Db) GetUsers() []User {
	rows, err := d.Conn.Query(context.Background(), "SELECT id, name, age, profession, friendly FROM users")
	if err != nil {
		log.Println(err)
	}
	return processRows(rows)
}

// GetUsersByName is called within our user query for graphql
func (d *Db) GetUsersByName(name string) []User {
	// Prepare query, takes a name argument, protects from sql injection
	rows, err := d.Conn.Query(context.Background(), "SELECT id, name, age, profession, friendly FROM users WHERE name=$1", name)
	if err != nil {
		log.Println(err)
	}
	return processRows(rows)
}
