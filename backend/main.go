package main

import (
	"backend/db"
	"backend/handlers"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	defer db.Db.Close()

	http.HandleFunc("/ping-table", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetPingResults(w, r)
		} else {
			http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handlers.PutPingResult(w, r)
		} else {
			http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		}
	})

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(http.DefaultServeMux)

	port := "8080"
	log.Println("Backend API running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
