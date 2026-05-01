package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// database details

const (
	postgres_host     = "db"
	postgres_port     = 5432
	postgres_user     = "postgres"
	postgres_password = "postgres"
	postgres_dbname   = "my_db"
)

// db variable to store the adderess of the our database
var Db *sql.DB

func init() {
	// create a connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)
	var err error
	// estbalish the connection to the database server using drive lib/pq
	Db, err = sql.Open("postgres", db_info)
	// handle error
	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully connected")
	}
}
