package handlers

import (
	"backend/db"
	"encoding/json"
	"log"
	"net/http"
)

func GetPingResults(w http.ResponseWriter, r *http.Request) {
	log.Println("Frontend asked for table")

	rows, err := db.Db.Query("SELECT * FROM Containers ORDER BY last_checked DESC")
	if err != nil {
		log.Println(err)
		http.Error(w, "Database request error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []db.PingResult
	for rows.Next() {
		var result db.PingResult
		if err := rows.Scan(&result.ContainerID, &result.IPAddress, &result.ResponseTime, &result.LastChecked); err != nil {
			log.Println(err)
			http.Error(w, "Data reading error", http.StatusInternalServerError)
			return
		}
		results = append(results, result)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func PutPingResult(w http.ResponseWriter, r *http.Request) {
	log.Println("Pinger sent ping info")

	var results []db.PingResult

	if err := json.NewDecoder(r.Body).Decode(&results); err != nil {
		log.Println(err)
		http.Error(w, "Parsing JSON error", http.StatusBadRequest)
		return
	}

	tx, err := db.Db.Begin()
	if err != nil {
		log.Println(err)
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	successRecord := 0

	for _, result := range results {
		_, err := tx.Exec(`
			INSERT INTO containers (container_id, ip_address, response_time, last_checked) 
			VALUES ($1, $2, $3, $4) 
			ON CONFLICT (container_id) 
			DO UPDATE SET 
				ip_address = EXCLUDED.ip_address, 
				response_time = EXCLUDED.response_time, 
				last_checked = EXCLUDED.last_checked;
		`, result.ContainerID, result.IPAddress, result.ResponseTime, result.LastChecked)

		if err != nil {
			log.Println(err)
			continue
		}

		successRecord++
	}

	if successRecord > 0 {
		if err := tx.Commit(); err != nil {
			log.Println(err)
			http.Error(w, "Transaction commit error", http.StatusInternalServerError)
			return
		}
	} else {
		tx.Rollback()
		http.Error(w, "No records were saved", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
