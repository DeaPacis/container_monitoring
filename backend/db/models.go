package db

type PingResult struct {
	ContainerID  string `json:"container_id"`
	IPAddress    string `json:"ip_address"`
	ResponseTime int    `json:"response_time"`
	LastChecked  string `json:"last_checked"`
}
