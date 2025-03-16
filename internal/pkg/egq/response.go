package egq

import "time"

type ResponseMeta struct {
	Timestamp time.Time `json:"timestamp"`
}

type Response struct {
	Meta ResponseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}
