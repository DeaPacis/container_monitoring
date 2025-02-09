package main

import (
	"log"
	"pinger/api"
	"pinger/containers"
	"pinger/models"
	"pinger/utils"
	"sync"
	"time"
)

func main() {
	scrapeInterval := utils.GetScrapeInterval()

	for {
		containersList, err := containers.GetContainers()
		if err != nil {
			log.Println(err)
			time.Sleep(scrapeInterval)
			continue
		}

		var results []models.PingResult
		var mu sync.Mutex
		var wg sync.WaitGroup

		for _, container := range containersList {
			wg.Add(1)

			go func(container models.PingResult) {
				defer wg.Done()

				responseTime, err := containers.PingContainer(container.IPAddress)
				if err != nil {
					log.Printf("Ping error %s (%s): %v\n", container.ContainerID, container.IPAddress, err)
					return
				}

				container.ResponseTime = responseTime
				container.LastChecked = time.Now().Format("2006-01-02 15:04:05")

				mu.Lock()
				results = append(results, container)
				mu.Unlock()
			}(container)
		}

		wg.Wait()

		if len(results) > 0 {
			api.SendResult(results)
		}

		time.Sleep(scrapeInterval)
	}
}
