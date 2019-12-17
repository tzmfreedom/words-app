package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		panic(err)
	}
}
