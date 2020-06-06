package utils

import (
	"log"
	"os"
)

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicln("ENV missing, key: " + k)
	}
	return v
}
