package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/", WithBasicAuth(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./public")).ServeHTTP(w, r)
	}))
	idSentenceRegexp := regexp.MustCompile(`^/sentences/(\d+)$`)
	http.HandleFunc("/sentences/", WithBasicAuth(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			if !idSentenceRegexp.MatchString(r.RequestURI) {
				log.Fatal(err)
				return
			}
			matches := idSentenceRegexp.FindStringSubmatch(r.RequestURI)

			var req UpdateSentenceRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				log.Fatal(err)
				return
			}
			id, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = updateSentence(db, int64(id), req.Value)
			if err != nil {
				log.Fatal(err)
				return
			}
			res, err := json.Marshal(&UpdateSentenceResponse{Id: id})
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Fprint(w, string(res))
		case http.MethodDelete:
			if !idSentenceRegexp.MatchString(r.RequestURI) {
				log.Fatal(err)
				return
			}
			matches := idSentenceRegexp.FindStringSubmatch(r.RequestURI)
			id, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = deleteSentence(db, int64(id))
			if err != nil {
				log.Fatal(err)
				return
			}
			res, err := json.Marshal(&DeleteSentenceResponse{Id: id})
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Fprint(w, string(res))
		}
	}))
	http.HandleFunc("/sentences", WithBasicAuth(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var req CreateSentenceRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = createSentence(db, req.Value)
			if err != nil {
				log.Fatal(err)
				return
			}
			res, err := json.Marshal(&CreateSentenceResponse{Id: 0})
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Fprint(w, string(res))
		case http.MethodGet:
			sentences, err := findAllSentence(db)
			if err != nil {
				log.Fatal(err)
				return
			}
			res, err := json.Marshal(&SentencesResponse{Records: sentences})
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Fprint(w, string(res))
		}
	}))
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func WithBasicAuth(f func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	basicAuthUser := os.Getenv("USER")
	basicAuthPass := os.Getenv("PASS")
	return func(w http.ResponseWriter, r *http.Request) {
		if user, pass, ok := r.BasicAuth(); !ok || user != basicAuthUser || pass != basicAuthPass {
			w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Not authorized", 401)
			return
		}
		f(w, r)
	}
}

func debug(args ...interface{}) {
	pp.Println(args...)
}
