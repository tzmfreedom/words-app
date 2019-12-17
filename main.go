package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	port := os.Getenv("PORT")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/sentences", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, value FROM sentences")
		if err != nil {
			panic(err.Error())
		}

		columns, err := rows.Columns()
		if err != nil {
			panic(err.Error())
		}
		fmt.Fprintln(w, columns[0])
		for rows.Next() {
			var id int
			var value string
			rows.Scan(&id, &value)
			fmt.Fprintf(w, "%d => %s\n", id, value)
		}
	})
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
