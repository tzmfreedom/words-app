package main

import (
	"database/sql"
	"errors"
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
		var value, createdAt, updatedAt string
		err := rows.Scan(&id, &value, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		sentences = append(sentences, &Sentence{
			Id:        id,
			Value:     value,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return sentences, nil
}

func findSentence(db *sql.DB, id int) (*Sentence, error) {
	rows, err := db.Query("SELECT id, value, created_at, updated_at FROM sentences WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var id int
		var value, createdAt, updatedAt string
		err := rows.Scan(&id, &value, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		return &Sentence{
			Id:        id,
			Value:     value,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, nil
	}
	return nil, errors.New("not found")
}

func createSentence(db *sql.DB, value string) (int, error) {
	row := db.QueryRow("INSERT INTO sentences(value, created_at, updated_at) VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id;", value)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func updateSentence(db *sql.DB, id int, value string) error {
	_, err := db.Exec("UPDATE sentences SET value = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", value, id)
	return err
}

func deleteSentence(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM sentences WHERE id = $1", id)
	return err
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

type ErrorResponse struct {
	Error string `json:"error"`
}
