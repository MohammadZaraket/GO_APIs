package main

/*import (
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

/*type Tag struct {
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

}*/
/*
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	ID   int    `json:"id"`
	Name string `json:"first_name"`
}

var Articles = []Article{}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:@/facebookdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT id, first_name FROM users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Article

		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		Articles = append(Articles, tag)

	}

	handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

*/
