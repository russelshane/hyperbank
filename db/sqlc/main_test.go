package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/russelshane/hyperbank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error while connecting to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}