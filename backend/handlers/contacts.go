package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"backend/database"
	"backend/models"
)

var contacts = []models.Contact{}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, name, phone, relation FROM contacts
	`)
	if err != nil {
		http.Error(w, "Database error ", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var record models.Contact
		err := rows.Scan(&record.ID, &record.Name, &record.Phone, &record.Relation)
		if err != nil {
			http.Error(w, "Error reading row", http.StatusInternalServerError)
			return
		}
		contacts = append(contacts, record)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	contacts = append(contacts, c)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}
