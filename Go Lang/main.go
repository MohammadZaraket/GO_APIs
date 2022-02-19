package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//fmt.Println(testapi.Test())
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:@/facebookdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO blocked VALUES (5,2,3)")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}
