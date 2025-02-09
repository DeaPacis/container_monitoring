package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"pinger/models"
)

func SendResult(results []models.PingResult) {
	backendURL := os.Getenv("BACKEND_URL")

	data, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
		return
	}

	request, err := http.NewRequest(http.MethodPut, backendURL, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println(response.Status)
	} else {
		log.Printf("Data sent: %+v\n", results)
	}
}
