package containers

import (
	"log"
	"os/exec"
	"pinger/models"
	"strings"
)

func GetContainers() ([]models.PingResult, error) {
	cmd := exec.Command("docker", "ps", "-q")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	containerIDs := strings.Fields(string(output))
	var results []models.PingResult

	for _, id := range containerIDs {
		cmd := exec.Command("docker", "inspect", "-f", "{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}", id)
		ipOutput, err := cmd.Output()
		if err != nil {
			log.Println(err)
			continue
		}

		ip := strings.TrimSpace(string(ipOutput))
		if ip == "" {
			log.Printf("Container %s has no IP\n", id)
			continue
		}

		results = append(results, models.PingResult{ContainerID: id, IPAddress: ip})
	}

	return results, nil
}
