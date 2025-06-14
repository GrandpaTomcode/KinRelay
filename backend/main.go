package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"backend/database"
	"backend/handlers"
	"backend/middleware"

	"github.com/gorilla/mux"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    dbURL := os.Getenv("KINRELAY_DB_URL")
    log.Println("DB URL: ", dbURL)
    database.Init()

    

    r := mux.NewRouter()
    r.HandleFunc("/api/ping-db", func (w http.ResponseWriter, r *http.Request)  {
        err := database.Pool.Ping(r.Context())
        if err != nil {
            http.Error(w, "Database connection failed", http.StatusInternalServerError)
            return
        }
        w.Write([]byte("Database OK"))
    }).Methods("GET")

    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.RequirePIN)

    api.HandleFunc("/contacts", handlers.GetContacts).Methods("GET")
    api.HandleFunc("/contacts", handlers.CreateContact).Methods("POST")

    log.Println("KinRelay Api running on: 9000")

    handler := middleware.EnableCORS(r)
    log.Fatal(http.ListenAndServe(":9000", handler))
}
