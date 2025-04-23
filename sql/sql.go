package sql

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Sql interface {
	Setup() error
	Shutdown() error
}

type sql struct {
	db         *sqlx.DB
	migrations string
	username   string
	password   string
	host       string
	port       int
	dbname     string
}

func NewSql(config Config) (Sql, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.Dbname))
	if err != nil {
		return nil, err
	}
	return sql{
		db:         db,
		migrations: config.Migrations,
		username:   config.Username,
		password:   config.Password,
		host:       config.Host,
		port:       config.Port,
		dbname:     config.Dbname,
	}, nil
}

func (m sql) Setup() error {
	migrations, err := migrate.New(fmt.Sprintf("file://%s", m.migrations), fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", m.username, m.password, m.host, m.port, m.dbname))

	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil {
		return err
	}
	fmt.Println("Migrations up")

	return nil
}

func (m sql) Shutdown() error {
	return m.db.Close()
}
