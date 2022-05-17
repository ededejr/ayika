package ayika

import (
	"io"
	"log"

	"github.com/joho/godotenv"
)

type load struct{}

// Load environment variables from .env file
func (l *load) Default() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err, "Could not load env")
	}
}

// Load environment variables from an array of .env files.
func (l *load) Local(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal(err, "Could not load env")
	}
}

// Load environment variables using an io.Reader, this
// is useful for loading env variables from a remote source.
func (l *load) WithReader(r io.Reader) map[string]string {
	envMap, err := godotenv.Parse(r)
	if err != nil {
		log.Fatal(err, "Could not load env")
	}
	return envMap
}

// Load environment variables into an in memory map.
func (l *load) IntoVariable() map[string]string {
	envMap, err := godotenv.Read()
	if err != nil {
		log.Fatal(err, "Could not load env")
	}
	return envMap
}
