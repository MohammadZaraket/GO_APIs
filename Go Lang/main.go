package main

import (
	"fmt"
	"log"

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

	/*insert, err := db.Query("INSERT INTO blocked VALUES (NULL,2,1)")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Success")*/

	type Tag struct {
		ID   int    `json:"id"`
		Name string `json:"first_name"`
	}

	results, err := db.Query("SELECT id, first_name FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}

}
