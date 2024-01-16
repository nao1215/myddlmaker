package main

import (
	"log"

	"github.com/shogo82148/myddlmaker"
	schema "github.com/shogo82148/myddlmaker/testdata/issue67"
)

func main() {
	m, err := myddlmaker.New(&myddlmaker.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m.AddStructs(&schema.User{})

	if err := m.GenerateFile(); err != nil {
		log.Fatal(err)
	}
	if err := m.GenerateGoFile(); err != nil {
		log.Fatal(err)
	}
}
