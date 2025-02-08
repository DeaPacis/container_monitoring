package main

import (
	"backend/db"
	"backend/handlers"
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

	port := "8080"
	log.Println("Backend API running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
