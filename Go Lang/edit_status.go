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

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		post_id := r.Form.Get("post_id")
		status := r.Form.Get("status")
		fmt.Println("Go MySQL Tutorial")
		db, err := sql.Open("mysql", "root:@/facebookdb")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		insert, err := db.Query("UPDATE posts SET status=? WHERE id=?;", status, post_id)
		if err != nil {
			panic(err.Error())
		} else {
			resp := make(map[string]string)
			resp["status"] = "Post Edited!"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			defer insert.Close()
			w.Write(jsonResp)
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
	mux.HandleFunc("/post", PostHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}
