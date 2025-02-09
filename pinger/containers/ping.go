package containers

import (
	"github.com/go-ping/ping"
	"time"
)

func PingContainer(ip string) (int, error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return 0, err
	}

	pinger.Count = 3
	pinger.Timeout = time.Second * 5

	err = pinger.Run()
	if err != nil {
		return 0, err
	}

	stats := pinger.Statistics()
	return int(stats.MaxRtt.Microseconds()), nil
}
