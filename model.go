package main

import (
	"database/sql"
)

type Sentence struct {
	Id        int    `json:"id"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func findAllSentence(db *sql.DB) ([]*Sentence, error) {
	rows, err := db.Query("SELECT id, value, created_at, updated_at FROM sentences")
	if err != nil {
		return nil, err
	}

	sentences := []*Sentence{}
	for rows.Next() {
		var id int
		var value string
		rows.Scan(&id, &value)
		sentences = append(sentences, &Sentence{
			Id:    id,
			Value: value,
		})
	}
	return sentences, nil
}

func createSentence(db *sql.DB, value string) (int64, error) {
	_, err := db.Exec("INSERT INTO sentences(value, created_at, updated_at) VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);", value)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func updateSentence(db *sql.DB, id int64, value string) (int64, error) {
	_, err := db.Exec("UPDATE sentences SET value = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", value, id)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func deleteSentence(db *sql.DB, id int64) (int64, error) {
	_, err := db.Exec("DELETE FROM sentences WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

type SentencesResponse struct {
	Records []*Sentence `json:"records"`
}

type CreateSentenceRequest struct {
	Value string `json:"value"`
}

type CreateSentenceResponse struct {
	Id int `json:"id"`
}

type UpdateSentenceRequest struct {
	Value string `json:"value"`
}

type UpdateSentenceResponse struct {
	Id int `json:"id"`
}

type DeleteSentenceResponse struct {
	Id int `json:"id"`
}
