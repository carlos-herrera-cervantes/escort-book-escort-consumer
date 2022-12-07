package config

import (
	"os"
	"sync"
)

type postgres struct {
	Databases postgresDatabases
}

type postgresDatabases struct {
	EscortProfile string
}

var singlePostgres *postgres
var lock = &sync.Mutex{}

func InitializePostgres() *postgres {
	if singlePostgres != nil {
		return singlePostgres
	}

	lock.Lock()
	defer lock.Unlock()

	singlePostgres = &postgres{
		Databases: postgresDatabases{
			EscortProfile: os.Getenv("DATABASE_URI"),
		},
	}

	return singlePostgres
}
