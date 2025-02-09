package containers

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"log"
	"pinger/models"

	"github.com/docker/docker/client"
)

func GetContainers() ([]models.PingResult, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var results []models.PingResult

	for _, c := range containers {
		inspect, err := cli.ContainerInspect(context.Background(), c.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		var ipAddress string
		for _, network := range inspect.NetworkSettings.Networks {
			ipAddress = network.IPAddress
			break
		}
		if ipAddress == "" {
			continue
		}

		results = append(results, models.PingResult{
			ContainerID: c.ID[:12],
			IPAddress:   ipAddress,
		})
	}

	return results, nil
}
