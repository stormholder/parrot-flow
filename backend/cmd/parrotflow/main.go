package main

import (
	"flag"
	"parrotflow/cmd/parrotflow/migrations"
)

var (
	dbPath string
)

func main() {
	flag.StringVar(&dbPath, "database", "store.db", "Database file path")
	flag.Parse()

	migrations.Init(dbPath)
}
