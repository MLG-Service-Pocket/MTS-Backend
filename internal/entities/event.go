package entities

import "time"

type Event struct {
	Title            string    `json:"title"`
	Category         string    `json:"category"`
	Location         string    `json:"event_geo"`
	MinPrice         int       `json:"eventMinPrice"`
	ClosestStartTime time.Time `json:"eventClosestDateTime"`
	LastStartTime    time.Time `json:"lastEventDateTime"`
	Tags             []string  `json:"tags"`
}
