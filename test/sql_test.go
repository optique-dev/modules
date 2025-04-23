package test

import (
	"os"
	"testing"

	"github.com/optique-dev/modules/sql"
)

func InitSQL() sql.Sql {
	current_dir, _ := os.Getwd()
	migrations := current_dir + "/migrations"
	app, err := sql.NewSql(sql.Config{
		Migrations: migrations,
		Username:   "optique",
		Password:   "optique",
		Host:       "0.0.0.0",
		Port:       5432,
		Dbname:     "optique",
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
