package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func wrap(f func(db *sql.DB)) {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseUrl)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	db.Exec("TRUNCATE TABLE sentences")
	f(db)
}

func createTestSentence(db *sql.DB, v string) *Sentence {
	id, err := createSentence(db, v)
	if err != nil {
		panic(err)
	}
	sentence, err := findSentence(db, id)
	if err != nil {
		panic(err)
	}
	return sentence
}

func testHandle(db *sql.DB, method, target, body string) *httptest.ResponseRecorder {
	handler := router(db)
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	basicAuthUser := os.Getenv("USER")
	basicAuthPass := os.Getenv("PASS")
	req.SetBasicAuth(basicAuthUser, basicAuthPass)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func TestGetSentences(t *testing.T) {
	wrap(func(db *sql.DB) {
		createTestSentence(db, "foo")
		createTestSentence(db, "bar")

		rec := testHandle(db, http.MethodGet, "/api/sentences", "")

		assert.Equal(t, http.StatusOK, rec.Code)
		var r SentencesResponse
		err := json.Unmarshal(rec.Body.Bytes(), &r)
		if err != nil {
			t.Fatal(err.Error())
		}
		assert.Equal(t, 2, len(r.Records))
		assert.EqualValues(t, "foo", r.Records[0].Value)
		assert.EqualValues(t, "bar", r.Records[1].Value)
	})
}

func TestGetSentence(t *testing.T) {
	wrap(func(db *sql.DB) {
		sentence1 := createTestSentence(db, "foo")
		createTestSentence(db, "bar")

		rec := testHandle(db, http.MethodGet, fmt.Sprintf("/api/sentences/%d", sentence1.Id), "")

		assert.Equal(t, http.StatusOK, rec.Code)
		var s Sentence
		err := json.Unmarshal(rec.Body.Bytes(), &s)
		if err != nil {
			t.Fatal(err.Error())
		}
		assert.Equal(t, "foo", s.Value)
		assert.Equal(t, sentence1.Id, s.Id)
		assert.Equal(t, sentence1.CreatedAt, s.CreatedAt)
		assert.Equal(t, sentence1.UpdatedAt, s.UpdatedAt)
	})
}

func TestCreateSentence(t *testing.T) {
	wrap(func(db *sql.DB) {
		body, err := json.Marshal(CreateSentenceRequest{Value: "new"})
		if err != nil {
			t.Fatal(err.Error())
		}
		rec := testHandle(db, http.MethodPost, "/api/sentences", string(body))

		assert.Equal(t, http.StatusOK, rec.Code)
		var r CreateSentenceResponse
		err = json.Unmarshal(rec.Body.Bytes(), &r)
		if err != nil {
			t.Fatal(err.Error())
		}
		sentence, err := findSentence(db, r.Id)
		if err != nil {
			t.Fatal(err.Error())
		}
		assert.Equal(t, "new", sentence.Value)
	})
}

func TestUpdateSentence(t *testing.T) {
	wrap(func(db *sql.DB) {
		s := createTestSentence(db, "foo")

		body, err := json.Marshal(UpdateSentenceRequest{Value: "update"})
		if err != nil {
			t.Fatal(err.Error())
		}
		rec := testHandle(db, http.MethodPut, fmt.Sprintf("/api/sentences/%d", s.Id), string(body))

		assert.Equal(t, http.StatusOK, rec.Code)
		var r UpdateSentenceResponse
		err = json.Unmarshal(rec.Body.Bytes(), &r)
		if err != nil {
			t.Fatal(err.Error())
		}
		sentence, err := findSentence(db, r.Id)
		if err != nil {
			t.Fatal(err.Error())
		}
		assert.Equal(t, "update", sentence.Value)
	})
}

func TestDeleteSentence(t *testing.T) {
	wrap(func(db *sql.DB) {
		s := createTestSentence(db, "foo")

		rec := testHandle(db, http.MethodDelete, fmt.Sprintf("/api/sentences/%d", s.Id), "")

		assert.Equal(t, http.StatusOK, rec.Code)
		var r UpdateSentenceResponse
		err := json.Unmarshal(rec.Body.Bytes(), &r)
		if err != nil {
			t.Fatal(err.Error())
		}
		sentence, err := findSentence(db, r.Id)
		if sentence != nil {
			t.Fatal("sentence should be nil")
		}
	})
}
