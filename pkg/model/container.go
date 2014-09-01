package model

type Container struct {
	BuildID     int64     `meddler:"build_id,pk"       json:"build_id"`
	Containers  string    `meddler:"containers"        json:"containers"`
}
