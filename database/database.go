package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/infuman69/expense-tracker-api/config"
	_ "github.com/lib/pq"
)

var (
	host     = config.Config("HOST")
	port     = config.Config("PORT")
	username = config.Config("USERNAME_DB")
	password = config.Config("PASSWORD")
	dbname   = config.Config("DBNAME")
)

func Database() {
	port_conv ,err:= strconv.ParseInt(port, 10,32)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port_conv, username, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected Successful")
}
