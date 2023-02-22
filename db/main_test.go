package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	db "github.com/jbattistella/movieratings/db/sqlc"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var testQueries *db.Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql:///test_ratings?sslmode=disable"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("dropdb", "--if-exists", "test_ratings")
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to drop")
		log.Fatal(err)
	}
	cmd = exec.Command("createdb", "test_ratings")
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to create")
		log.Fatal(err)
	}

	mig, err := migrate.New(
		"file://migration",
		dbSource)
	if err != nil {
		log.Fatal(err)
	}
	if err := mig.Up(); err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())

}
