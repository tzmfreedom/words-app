package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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
		http.FileServer(http.Dir("./client/dist")).ServeHTTP(w, r)
	}))
	idSentenceRegexp := regexp.MustCompile(`^/sentences/(\d+)$`)
	http.HandleFunc("/sentences/", WithBasicAuth(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if !idSentenceRegexp.MatchString(r.RequestURI) {
				handleNotFound(w)
				return
			}
			matches := idSentenceRegexp.FindStringSubmatch(r.RequestURI)
			id, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			sentence, err := findSentence(db, id)
			if err != nil {
				handleNotFound(w)
				return
			}
			res, err := json.Marshal(sentence)
			if err != nil {
				panic(err)
			}
			fmt.Fprint(w, string(res))
		case http.MethodPut:
			if !idSentenceRegexp.MatchString(r.RequestURI) {
				handleNotFound(w)
				return
			}
			matches := idSentenceRegexp.FindStringSubmatch(r.RequestURI)

			var req UpdateSentenceRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				handleInvalidRequest(w)
				return
			}
			id, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			err = updateSentence(db, id, req.Value)
			if err != nil {
				log.Fatal(err)
				return
			}
			res, err := json.Marshal(&UpdateSentenceResponse{Id: id})
			if err != nil {
				panic(err)
			}
			fmt.Fprint(w, string(res))
		case http.MethodDelete:
			if !idSentenceRegexp.MatchString(r.RequestURI) {
				handleNotFound(w)
				return
			}
			matches := idSentenceRegexp.FindStringSubmatch(r.RequestURI)
			id, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			err = deleteSentence(db, id)
			if err != nil {
				log.Fatal(err)
				return
			}
			res, err := json.Marshal(&DeleteSentenceResponse{Id: id})
			if err != nil {
				panic(err)
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
				handleInvalidRequest(w)
				return
			}
			id, err := createSentence(db, req.Value)
			if err != nil {
				log.Fatal(err)
				return
			}
			res, err := json.Marshal(&CreateSentenceResponse{Id: id})
			if err != nil {
				panic(err)
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
				panic(err)
			}
			fmt.Fprint(w, string(res))
		}
	}))
	handler := cors.AllowAll().Handler(http.DefaultServeMux)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		panic(err)
	}
}

func handleInvalidRequest(w http.ResponseWriter) {
	handleErrorResponse(w, http.StatusBadRequest, "Invalid Request")
}

func handleNotFound(w http.ResponseWriter) {
	handleErrorResponse(w, http.StatusNotFound, "Not Found")
}

func handleErrorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	w.WriteHeader(statusCode)
	buf, err := json.Marshal(ErrorResponse{Error: errorString})
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(buf))
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
