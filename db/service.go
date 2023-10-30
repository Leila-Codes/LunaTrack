package db

import "Luna_Track/db/postgres"

type DatabaseType int

const (
	Postgres DatabaseType = iota
)

type DatabaseConfig struct {
	Host, User, Pass, Name string
	Port                   int
}

func ConnectDatabase(dbType DatabaseType, config DatabaseConfig) IDatabase {
	switch dbType {
	case Postgres:
		return postgres.NewPostgresDB(config.Host, config.User, config.Pass, config.Name, config.Port)
	default:
		panic("unsupported database impl type" + string(dbType))
	}
}
