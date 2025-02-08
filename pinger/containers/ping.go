package containers

import (
	"github.com/go-ping/ping"
	"time"
)

func PingContainer(ip string) (int, error) {
	pinger, err := ping.NewPinger(ip)
	//pinger, err := ping.NewPinger("1.1.1.1")
	if err != nil {
		return 0, err
	}

	pinger.Count = 3
	pinger.Timeout = time.Second * 5
	pinger.SetPrivileged(true) // Linux

	err = pinger.Run()
	if err != nil {
		return 0, err
	}

	stats := pinger.Statistics()
	//resp := float64(stats.MaxRtt.Seconds())
	//log.Println(resp)
	return int(stats.AvgRtt.Milliseconds()), nil
}
