package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	raspSql "github.com/mrtc0/gorasp/contrib/database/sql"
)

func main() {
	driverName, err := raspSql.Register("sqlite3")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(driverName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE users(id int, email text, password text);")

	ctx := context.Background()
	_, err = db.QueryContext(ctx, "SELECT * FROM users WHERE email = ? AND password = ?;", "user@example.com", "password")
	if err != nil {
		panic(err)
	}

	input := "user@example.com' OR 1=1 --"
	_, err = db.QueryContext(ctx, fmt.Sprintf(
		"SELECT * FROM users WHERE email = %s AND password = %s;", input, "",
	))
	if err != nil {
		fmt.Printf("SQLi detected: %v\n", err)
	}
}
