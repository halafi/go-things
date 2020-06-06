package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/halafi/go-things/go-graphql-api/pkg/utils"
	"github.com/jackc/pgx/v4"
)

// should be run with make from root - parent project dir
const rootDir = "db/seeds"

func loadSQLFiles() []string {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func main() {
	dbURL := utils.MustGet("DB_URL")
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	files := loadSQLFiles()

	// TODO: ensure sorting if needed
	for _, file := range files {
		sql, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Executing " + file)
		_, err = conn.Exec(context.Background(), string(sql))

	}
}
