package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	flagPort = flag.String("port", "9000", "Port to listen on")
)

var results []string

type Article struct {
	ID int `json:"id"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		//var Articles = []Article{}
		r.ParseForm()
		user_id := r.Form.Get("user_id")
		post_id := r.Form.Get("post_id")
		fmt.Println("Go MySQL Tutorial")
		db, err := sql.Open("mysql", "root:@/facebookdb")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		var enough bool
		if err := db.QueryRow("SELECT id FROM likes WHERE post_id=? AND user_id=?", post_id, user_id).Scan(&enough); err != nil {
			if err == sql.ErrNoRows {

				insert, err := db.Query("INSERT INTO likes VALUES (NULL,?,?)", post_id, user_id)
				if err != nil {
					panic(err.Error())

				}
				defer insert.Close()
				fmt.Println("sucess")

			} else {

				delete, err := db.Query("DELETE FROM likes WHERE post_id=? AND user_id=?", post_id, user_id)
				if err != nil {
					panic(err.Error())

				}

				defer delete.Close()
				fmt.Println("sucess")
			}

		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
}

func main() {

	results = append(results, time.Now().Format(time.RFC3339))
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHandler)
	mux.HandleFunc("/post", PostHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))

}
