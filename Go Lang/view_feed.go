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
	ID     int    `json:"id"`
	STATUS string `json:"status"`
	LIKES  string `json:"number_likes"`
	FNAME  string `json:"first_name"`
	LNAME  string `json:"last_name"`
	DATE   string `json:"date_time"`
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
		var Articles = []Article{}
		r.ParseForm()
		user_id := r.Form.Get("user_id")
		fmt.Println("Go MySQL Tutorial")
		db, err := sql.Open("mysql", "root:@/facebookdb")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		results, err := db.Query("SELECT posts.id,STATUS,number_likes,first_name,last_name,date_time FROM posts,users WHERE user_id IN (SELECT friend_id from friends where user_id= ? and friend_id not in (SELECT user_id from blocked where blocked_id= ?)) and posts.user_id= users.id ORDER by date_time DESC", user_id, user_id)
		if err != nil {
			panic(err.Error())
		}

		for results.Next() {
			var tag Article
			err = results.Scan(&tag.ID, &tag.STATUS, &tag.LIKES, &tag.FNAME, &tag.LNAME, &tag.DATE)
			if err != nil {
				panic(err.Error())
			}
			Articles = append(Articles, tag)
		}
		json.NewEncoder(w).Encode(Articles)

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
