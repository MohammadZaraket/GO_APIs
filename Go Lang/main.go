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

	insert, err := db.Query("INSERT INTO blocked VALUES (NULL,2,1)")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Success")

}
