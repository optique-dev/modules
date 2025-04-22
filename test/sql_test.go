package test

import (
	"testing"

	"github.com/Courtcircuits/optique-modules/sql"
)

func InitSQL() sql.Sql {
	app, err := sql.NewSql(sql.Config{
		Migrations: "./migrations",
		Username:   "optique",
		Password:   "optique",
		Host:       "0.0.0.0",
		Port:       5432,
		Dbname:   "optique",
	})
	if err != nil {
		panic(err)
	}
	return app
}

func SetupSQL(t *testing.T, app sql.Sql) {
	if err := app.Setup(); err != nil {
		t.Fatal(err)
	}
}

func TestSQLMigrate(t *testing.T) {
	app := InitSQL()
	defer app.Shutdown()
	SetupSQL(t, app) // this will run the migrations
}
