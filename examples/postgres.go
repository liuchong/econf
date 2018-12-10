package examples

import (
	"fmt"
)

type postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func (pg *postgres) URI() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		pg.User, pg.Password, pg.Host, pg.Port, pg.Dbname,
	)
}

// Postgres is configuration of postgresql database connection
var Postgres = postgres{
	Host:     "postgres",
	Port:     "5432",
	User:     "postgres",
	Password: "123456",
	Dbname:   "postgres",
}
