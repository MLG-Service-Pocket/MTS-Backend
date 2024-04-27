package entities

type Event struct {
	Title            string   `json:"title"`
	Category         string   `json:"category"`
	Location         string   `json:"event_geo"`
	MinPrice         int      `json:"eventMinPrice"`
	ClosestStartTime string   `json:"eventClosestDateTime"`
	LastStartTime    string   `json:"lastEventDateTime"`
	Tags             []string `json:"tags"`
}
